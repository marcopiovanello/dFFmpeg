package ffmpeg

import (
	"context"
	"os"

	"github.com/marcopeocchi/sanji/internal/processor"
)

type Job struct {
	Context        context.Context
	Processor      processor.VideoProcessor
	OutputFile     string
	FileDescriptor *os.File
}
