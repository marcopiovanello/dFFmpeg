package ffmpeg

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/marcopeocchi/sanji/internal/ffmpeg/pb"
	"github.com/marcopeocchi/sanji/internal/ffprobe"
	"github.com/marcopeocchi/sanji/internal/processor"
)

type FFmpegServer struct {
	pb.UnimplementedFFmpegServer

	prepared sync.Map
	progress sync.Map
	logs     sync.Map

	metadata map[string]*ffprobe.FFprobeOutput
}

func NewFFmpegServer() pb.FFmpegServer {
	return &FFmpegServer{
		metadata: make(map[string]*ffprobe.FFprobeOutput),
	}
}

// PrepareConversion implements pb.FFmpegServer.
func (f *FFmpegServer) PrepareConversion(ctx context.Context, req *pb.PrepareConversionRequest) (*pb.PrepareConversionResponse, error) {
	outFile := fmt.Sprintf("%s%s", req.Id, filepath.Ext(req.Filename))

	fd, err := os.Create(outFile)
	if err != nil {
		return nil, err
	}

	log.Printf("prepared file %s\n", outFile)

	f.prepared.Store(req.Id, Job{
		Context:        context.Background(),
		Processor:      processor.NewFactory(processor.Encoder(req.Processor), int(req.Quality)),
		OutputFile:     outFile,
		FileDescriptor: fd,
	})

	return &pb.PrepareConversionResponse{
		Id:                     req.Id,
		Processor:              req.Processor,
		AdditionalFfmpegParams: req.AdditionalFfmpegParams,
		TempFilename:           outFile,
	}, nil
}

// StartConversion implements pb.FFmpegServer.
func (f *FFmpegServer) StartConversion(stream pb.FFmpeg_StartConversionServer) error {
	startTime := time.Now()

	chunk, err := stream.Recv()
	if err == io.EOF || err != nil {
		return nil
	}

	job, ok := f.prepared.Load(chunk.Id)
	if !ok {
		return errors.New("conversion not prepared")
	}

	var (
		ctx       = job.(Job).Context
		fd        = job.(Job).FileDescriptor
		processor = job.(Job).Processor
		inputFile = job.(Job).OutputFile

		// send all errors from goroutines spawned below
		errChan = make(chan error)
	)

	// copy the first chunk
	written, err := io.Copy(fd, bytes.NewReader(chunk.FileContent))
	if err != nil {
		return err
	}

	log.Printf("Received first %d bytes of %s\n", written, chunk.Id)

	// save the current file metadata, useful for determinate progress

	metadata, err := ffprobe.ParseFile(ctx, inputFile)
	if err != nil {
		log.Fatalln(err)
		errChan <- err
	}

	f.metadata[chunk.Id] = metadata

	// start converting with ffmpeg
	ffmpegOutput, err := processor.Process(ctx, inputFile)
	if err != nil {
		return err
	}

	log.Printf("Starting conversion with profile %d\n", processor)

	// continue copying and piping to the file
	go func() {
		for {
			chunk, err := stream.Recv()
			if err != nil {
				errChan <- err
				return
			}

			_, err = io.Copy(fd, bytes.NewReader(chunk.FileContent))
			if err != nil {
				break
			}
		}
	}()

	for {
		select {
		case event := <-ffmpegOutput:
			//TODO: refactor out to dedicated functions
			stream.Send(&pb.ConversionResponse{
				FfmpegOutput: event,
				ElapsedTime:  int32(time.Since(startTime)),
			})

			f.logs.Store(chunk.Id, event)

			metadata := f.metadata[chunk.Id]

			if metadata != nil {
				totalFrames, err := ffprobe.TotalFrames(&f.metadata[chunk.Id].Streams)
				if err != nil {
					log.Println("failed to get totalframes", err)
					errChan <- err
				}

				progress, err := parseProgress(string(event), totalFrames)
				if err == nil {
					f.progress.Store(chunk.Id, progress)
				}
			}
		case err := <-errChan:
			return err
		case <-ctx.Done():
			return ctx.Err()
			// case <-stream.Context().Done():
			// 	return stream.Context().Err()
		}
	}
}

// StopConversion implements pb.FFmpegServer.
func (f *FFmpegServer) StopConversion(ctx context.Context, req *pb.Query) (*pb.Query, error) {
	job, ok := f.prepared.Load(req.Id)
	if !ok {
		return nil, errors.New("conversion not prepared")
	}

	_, cancel := context.WithCancel(job.(Job).Context)
	cancel()

	f.prepared.Delete(req.Id)
	delete(f.metadata, req.Id)

	return &pb.Query{
		Id: req.Id,
	}, nil
}

// GetProgress implements pb.FFmpegServer.
func (f *FFmpegServer) GetProgress(req *pb.Query, stream pb.FFmpeg_GetProgressServer) error {
	var (
		ticker  = time.NewTicker(time.Second)
		errChan = make(chan error)
	)

	sendProgress := func() {
		v, ok := f.progress.Load(req.Id)
		if !ok {
			errChan <- errors.New("job " + req.Id + " has an unknown state")
		}

		progress := v.(*Progress)

		stream.Send(&pb.Progress{
			Id:      req.Id,
			BitRate: progress.BitRate,
			Ratio:   progress.Ratio,
			Fps:     int32(progress.FPS),
			Q:       int32(progress.Q),
		})
	}

	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case <-ticker.C:
			sendProgress()
		case err := <-errChan:
			return err
		}
	}
}
