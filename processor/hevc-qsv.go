package processor

import (
	"context"
	"os"
	"os/exec"
	"path"
	"strconv"
)

type HEVCQSVProcessor struct {
	ffmpegPath string
	cqp        int
}

func NewHEVCQSVProcessor(ffmpegPath string, cqp int) *HEVCQSVProcessor {
	return &HEVCQSVProcessor{
		ffmpegPath: ffmpegPath,
		cqp:        cqp,
	}
}

func (p *HEVCQSVProcessor) Process(ctx context.Context, input string) error {
	tempFile := path.Join(path.Dir(input), "."+path.Base(input))

	cqpString := strconv.Itoa(p.cqp)

	// Spawn a new ffmpeg process and convert a video to HEVC-10bit with the QSV
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
		"-q", cqpString,
		tempFile,
	)

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
