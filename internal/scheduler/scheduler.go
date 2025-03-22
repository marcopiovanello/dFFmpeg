package scheduler

import "context"

type Scheduler interface {
	Schedule(ctx context.Context, j ConversionJob)
	Pending(ctx context.Context) *[]ConversionJob
}
