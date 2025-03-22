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

type AV1Rav1eProcessor struct {
	ffmpegPath  string
	videoPreset int
}

func NewRav1eAV1Processor(path string, preset int) VideoProcessor {
	return &AV1Rav1eProcessor{
		ffmpegPath:  path,
		videoPreset: preset,
	}
}

func (p *AV1Rav1eProcessor) Process(ctx context.Context, input string) (<-chan []byte, error) {
	ffmpegOutput := make(chan []byte)

	if p.videoPreset < 1 {
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
		"-c:v", "librav1e",
		"-pix_fmt", "yuv420p10le",
		"-crf", "22",
		"-preset", strconv.Itoa(p.videoPreset),
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
