package processor

import (
	"bufio"
	"context"
	"errors"
	"log"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
	"github.com/marcopeocchi/sanji/internal/utils"
)

type HEVCVideoToolboxProcessor struct {
	ffmpegPath    string
	qualityPreset *QualityPreset
}

func NewHEVCVideoToolboxProcessor(path string, qp *QualityPreset) VideoProcessor {
	if qp == nil {
		qp = &QualityPreset{
			Quality: 65,
		}
	}

	return &HEVCVideoToolboxProcessor{
		ffmpegPath:    path,
		qualityPreset: qp,
	}
}

func (p *HEVCVideoToolboxProcessor) Process(ctx context.Context, input string) (<-chan []byte, error) {
	ffmpegOutput := make(chan []byte)

	if p.qualityPreset.Quality < 1 {
		return nil, errors.New("preset must be greater than zero")
	}

	tempFile := uuid.NewString() + filepath.Ext(input)

	// Spawn a new ffmpeg process and convert a video to AV1-10bit with the SVT
	// encoder, copying audio and subtitles streams.
	cmd := exec.CommandContext(ctx, p.ffmpegPath,
		"-i", input,
		"-map", "0",
		"-c:a", "copy",
		"-c:s", "copy",
		"-c:v", "hevc_videotoolbox",
		"-q:v", strconv.Itoa(p.qualityPreset.Quality),
		"-tag:v", "hvc1",
		"-pix_fmt", "p010le",
		tempFile,
	)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	go func() {
		scanner := bufio.NewScanner(stderr)
		scanner.Split(utils.FFMpegStdoutSplitFunc)
		for scanner.Scan() {
			ffmpegOutput <- scanner.Bytes()
		}
	}()
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			ffmpegOutput <- scanner.Bytes()
		}
	}()

	go func() {
		<-ctx.Done()

		// do something when asked to interrupt
		log.Println("stopping")
		close(ffmpegOutput)
	}()

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	return ffmpegOutput, nil
}
