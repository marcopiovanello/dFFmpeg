package scheduler

import (
	"context"
	"log/slog"

	"github.com/marcopeocchi/sanji/internal/processor"
)

type RoundRobin struct {
	concurrency int
	pipeline    chan struct{}
	proc        processor.VideoProcessor
	logger      *slog.Logger
}

func NewRoundRobin(c int, p processor.VideoProcessor, l *slog.Logger) Scheduler {
	return &RoundRobin{
		concurrency: c,
		pipeline:    make(chan struct{}, c),
		proc:        p,
		logger:      l,
	}
}

func (r *RoundRobin) Schedule(ctx context.Context, j ConversionJob) {
	go func() {
		r.pipeline <- struct{}{}

		r.logger.Info("starting", slog.String("file", j.InputFile))

		if _, err := r.proc.Process(ctx, j.InputFile); err != nil {
			r.logger.Error(
				"error while converting",
				slog.String("file", j.InputFile),
				slog.String("err", err.Error()),
			)
		}

		<-r.pipeline
	}()
}

func (r *RoundRobin) Pending(ctx context.Context) *[]ConversionJob {
	panic("unimplemented")
}
