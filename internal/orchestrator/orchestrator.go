package orchestrator

import (
	"context"

	"github.com/marcopeocchi/sanji/internal/ffmpeg/pb"
	"github.com/marcopeocchi/sanji/internal/processor"
)

type Orchestrator interface {
	StartJob(ctx context.Context, path string, encoder processor.Encoder, quality int) (string, error)
	StopJob(ctx context.Context, id string) error
	Aggregate(ctx context.Context) (<-chan *pb.Progress, error)
	Details(ctx context.Context, id string) (<-chan *pb.Progress, error)
}
