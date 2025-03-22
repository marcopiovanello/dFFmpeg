package ffprobe

import (
	"context"
	"encoding/json"
	"os/exec"
)

func ParseFile(ctx context.Context, path string) (*FFprobeOutput, error) {
	cmd := exec.CommandContext(ctx, "ffprobe",
		"-v", "quiet",
		"-print_format", "json",
		"-show_format",
		"-show_streams",
		path,
	)

	pipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer pipe.Close()

	cmd.Start()
	output := &FFprobeOutput{}

	if err := json.NewDecoder(pipe).Decode(output); err != nil {
		return nil, err
	}

	cmd.Wait()
	return output, nil
}
