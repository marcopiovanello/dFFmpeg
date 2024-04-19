package processor

import (
	"context"
	"os"
	"os/exec"
)

type AV1Processor struct {
	ffmpegPath  string
	videoPreset string
}

func NewAV1Processor(path, preset string) *AV1Processor {
	return &AV1Processor{
		ffmpegPath:  path,
		videoPreset: "6",
	}
}

func (p *AV1Processor) Process(ctx context.Context, input string) error {
	tempFile := "." + input

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
		"preset", p.videoPreset,
		tempFile,
	)

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
