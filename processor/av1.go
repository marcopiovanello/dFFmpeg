package processor

import (
	"bufio"
	"context"
	"log/slog"
	"os"
	"os/exec"
	"path"
)

type AV1Processor struct {
	ffmpegPath  string
	videoPreset string
	logger      *slog.Logger
}

func NewAV1Processor(path, preset string, logger *slog.Logger) *AV1Processor {
	return &AV1Processor{
		ffmpegPath:  path,
		videoPreset: "6",
		logger:      logger,
	}
}

func (p *AV1Processor) Process(ctx context.Context, input string) error {
	tempFile := path.Join(path.Dir(input), "."+path.Base(input))

	// Spawn a new ffmpeg process and convert a video to AV1-10bit with the SVT
	// encoder, copying audio and subtitles streams.
	cmd := exec.CommandContext(ctx, p.ffmpegPath,
		"-i", input,
		"-map", "0",
		"-c:a", "copy",
		"-c:s", "copy",
		"-c:v", "libsvtav1",
		"-pix_fmt", "yuv420p10le",
		"-crf", "22",
		"-preset", p.videoPreset,
		tempFile,
	)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	go func() {
		scanner := bufio.NewScanner(stderr)

		for scanner.Scan() {
			p.logger.Error(scanner.Text())
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stdout)

		for scanner.Scan() {
			p.logger.Info(scanner.Text())
		}
	}()

	if err := cmd.Start(); err != nil {
		return err
	}

	// Wait to conversion process to end if it errors abort and remove the
	// temp file
	if err := cmd.Wait(); err != nil {
		os.Remove(tempFile)
		return err
	}

	// Remove the temp file
	if err := os.Remove(input); err != nil {
		return err
	}

	// Move the converted video and rename it as the original
	if err := os.Rename(tempFile, input); err != nil {
		return err
	}

	return nil
}
