package processor

import "context"

type VideoProcessor interface {
	Process(ctx context.Context, input string) error
}
