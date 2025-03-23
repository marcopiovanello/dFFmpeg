package orchestrator

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
	"github.com/marcopeocchi/sanji/internal/ffmpeg/pb"
	loadbalancer "github.com/marcopeocchi/sanji/internal/load_balancer"
	"github.com/marcopeocchi/sanji/internal/processor"
	"github.com/marcopeocchi/sanji/internal/utils"
)

type GrpcOrchestrator struct {
	client   pb.FFmpegClient
	db       *memdb.MemDB
	requests chan loadbalancer.Request
}

func NewGrpcOrchestrator(client pb.FFmpegClient, nWorkers int) (Orchestrator, error) {
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		return nil, err
	}

	requests := make(chan loadbalancer.Request)

	lb := loadbalancer.New(nWorkers) //TODO: test with 8 cores
	go lb.Balance(requests)

	return &GrpcOrchestrator{
		client:   client,
		db:       db,
		requests: requests,
	}, nil
}

func (o *GrpcOrchestrator) Aggregate(ctx context.Context) (<-chan *pb.Progress, error) {
	tx := o.db.Txn(false)

	it, err := tx.Get("jobs", "nodes")
	if err != nil {
		return nil, err
	}

	fanInChannel := make(chan *pb.Progress)

	for obj := it.Next(); obj != nil; obj = it.Next() {
		j := obj.(*FFmpegJob)

		go func() {
			streamingClient, err := o.client.GetProgress(ctx, &pb.Query{Id: j.Id})
			if err != nil {
				log.Println(err)
				return
			}

			for {
				select {
				case <-ctx.Done():
					streamingClient.CloseSend()
					close(fanInChannel)
					return
				case <-streamingClient.Context().Done():
					return
				default:
					res, err := streamingClient.Recv()
					if err == io.EOF || err != nil {
						return
					}
					fanInChannel <- res
				}
			}
		}()
	}

	return fanInChannel, nil
}

// StartJob implements Orchestrator.
func (o *GrpcOrchestrator) StartJob(
	ctx context.Context,
	path string,
	encoder processor.Encoder,
	qp *processor.QualityPreset,
) (string, error) {
	id := uuid.NewString()

	var (
		crf     *int32
		preset  *int32
		quality *int32
	)

	if qp != nil {
		_crf := int32(qp.CRF)
		crf = &_crf

		_preset := int32(qp.Preset)
		preset = &_preset

		_quality := int32(qp.Quality)
		quality = &_quality
	}

	res, err := o.client.PrepareConversion(context.Background(), &pb.PrepareConversionRequest{
		Id:                     id,
		Filename:               path,
		Processor:              int32(encoder),
		Quality:                quality,
		Crf:                    crf,
		Preset:                 preset,
		AdditionalFfmpegParams: nil,
	})
	if err != nil {
		return "", err
	}

	tx := o.db.Txn(true)
	tx.Insert("jobs", FFmpegJob{
		Id:               res.Id,
		Node:             utils.GetLocalIP().String(),
		OriginalFilePath: path,
	})
	defer tx.Abort()
	tx.Commit()

	streamingClient, err := o.client.StartConversion(ctx)
	if err != nil {
		return "", err
	}

	fd, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer fd.Close()

	br := bufio.NewReader(fd)
	buf := make([]byte, int64(2048*1000))

	for {
		_, err := br.Read(buf)
		if err != nil {
			log.Println("succesfully sent file", err)
			break
		}

		streamingClient.Send(&pb.ConversionContent{
			Id:          id,
			FileContent: buf,
		})
	}

	o.requests <- loadbalancer.NewRequest(func() pb.FFmpeg_GetProgressClient {
		streamingClient, err := o.client.GetProgress(ctx, &pb.Query{Id: id})
		if err != nil {
			return nil
		}
		return streamingClient
	})

	buf = nil
	return id, nil
}

// StopJob implements Orchestrator.
func (o *GrpcOrchestrator) StopJob(ctx context.Context, id string) error {
	res, err := o.client.StopConversion(ctx, &pb.Query{
		Id: id,
	})
	if err != nil {
		return err
	}

	tx := o.db.Txn(true)
	tx.Delete("jobs", FFmpegJob{Id: res.Id})
	tx.Commit()

	return nil
}

// Details implements Orchestrator.
func (o *GrpcOrchestrator) Details(ctx context.Context, id string) (<-chan *pb.Progress, error) {
	streamingClient, err := o.client.GetProgress(ctx, &pb.Query{Id: id})
	if err != nil {
		return nil, err
	}

	progress := make(chan *pb.Progress)

	go func() {
		for {
			select {
			case <-ctx.Done():
				streamingClient.CloseSend()
				close(progress)
				return
			default:
				res, err := streamingClient.Recv()
				if err == io.EOF || err != nil {
					close(progress)
					return
				}
				progress <- res
			}
		}
	}()

	return progress, nil
}
