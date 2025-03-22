package processor

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/google/uuid"
)

type HEVCQSVProcessor struct {
	ffmpegPath string
	cqp        int
}

func NewHEVCQSVProcessor(ffmpegPath string, cqp int) VideoProcessor {
	return &HEVCQSVProcessor{
		ffmpegPath: ffmpegPath,
		cqp:        cqp,
	}
}

func (p *HEVCQSVProcessor) Process(ctx context.Context, input string) (<-chan []byte, error) {
	ffmpegOutput := make(chan []byte)

	if p.cqp < 1 {
		return nil, errors.New("constant quality profile must be greater than zero")
	}

	tempFile := fmt.Sprintf("%s.%s", uuid.NewString(), filepath.Ext(input))

	// Spawn a new ffmpeg process and convert a video to AV1-10bit with the SVT
	// encoder, copying audio and subtitles streams.
	cmd := exec.CommandContext(ctx, p.ffmpegPath,
		"-init_hw_device", "qsv=hw",
		"-filter_hw_device", "hw",
		"-i", input,
		"-map", "0",
		"-c:a", "copy",
		"-c:s", "copy",
		"-c:v", "hevc_qsv",
		"-pix_fmt", "p010le",
		"-profile:v", "main10",
		"-q", strconv.Itoa(p.cqp),
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
		close(ffmpegOutput)
	}()

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	return ffmpegOutput, nil
}
