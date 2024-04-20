package main

import (
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/marcopeocchi/sanji/config"
	"github.com/marcopeocchi/sanji/logging"
	"github.com/marcopeocchi/sanji/processor"
	"github.com/marcopeocchi/sanji/scheduler"
	"github.com/rjeczalik/notify"
)

func main() {
	config.LoadFile("config.yaml")

	logWriters := []io.Writer{
		os.Stdout,
		logging.NewObservableLogger(),
	}

	logger := slog.New(
		slog.NewTextHandler(io.MultiWriter(logWriters...), &slog.HandlerOptions{}),
	)

	eventChan := make(chan notify.EventInfo, 1)

	if err := notify.Watch(
		config.Instance().Root,
		eventChan,
		notify.All,
	); err != nil {
		log.Fatal(err)
	}

	defer notify.Stop(eventChan)

	var (
		p = processor.NewFactory(processor.SVT_AV1, logger)
		s = scheduler.NewRoundRobinScheduler(1, p, logger)
	)

	for event := range eventChan {
		logger.Info(event.Event().String())

		if event.Event() != notify.Write && event.Event() != notify.Rename {
			continue
		}

		filename := filepath.Base(event.Path())

		if strings.HasPrefix(filename, config.Instance().ReleasePrefix) &&
			!strings.HasPrefix(filename, ".") {
			s.Schedule(scheduler.ConversionJob{
				InputFile: event.Path(),
			})
		}
	}
}
