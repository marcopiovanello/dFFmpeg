package scheduler

import (
	"context"
	"log/slog"

	"github.com/marcopeocchi/sanji/processor"
)

type RoundRobinScheduler struct {
	concurrency int
	pipeline    chan struct{}
	proc        processor.VideoProcessor
	logger      *slog.Logger
}

func NewRoundRobinScheduler(c int, p processor.VideoProcessor, l *slog.Logger) *RoundRobinScheduler {
	return &RoundRobinScheduler{
		concurrency: c,
		pipeline:    make(chan struct{}, c),
		proc:        p,
		logger:      l,
	}
}

func (r *RoundRobinScheduler) Schedule(j ConversionJob) {
	go func() {
		r.pipeline <- struct{}{}

		r.logger.Info("starting", slog.String("file", j.InputFile))

		err := r.proc.Process(context.Background(), j.InputFile)
		if err != nil {
			r.logger.Error(
				"error while converting",
				slog.String("file", j.InputFile),
				slog.String("err", err.Error()),
			)
		}

		<-r.pipeline
	}()
}
