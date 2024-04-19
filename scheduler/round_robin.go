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
}

func NewRoundRobinScheduler(c int, p processor.VideoProcessor) *RoundRobinScheduler {
	return &RoundRobinScheduler{
		concurrency: c,
		pipeline:    make(chan struct{}, c),
		proc:        p,
	}
}

func (r *RoundRobinScheduler) Schedule(j ConversionJob) {
	go func() {
		r.pipeline <- struct{}{}
		slog.Info("starting", slog.String("file", j.InputFile))
		err := r.proc.Process(context.Background(), j.InputFile)
		if err != nil {
			slog.Error(
				"error while converting",
				slog.String("file", j.InputFile),
				slog.String("err", err.Error()),
			)
		}
		<-r.pipeline
	}()
}
