package main

import (
	"context"
	"flag"
	"io"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/marcopeocchi/sanji/config"
	"github.com/marcopeocchi/sanji/logging"
	"github.com/marcopeocchi/sanji/processor"
	"github.com/marcopeocchi/sanji/scheduler"
	"github.com/rjeczalik/notify"
)

var configPath string

func main() {
	flag.StringVar(&configPath, "c", "config.yaml", "full path of config file")
	flag.Parse()

	config.LoadFile(configPath)

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

	ctx, _ := signal.NotifyContext(context.Background(),
		os.Interrupt,
		os.Kill,
		syscall.SIGTERM,
		syscall.SIGKILL,
	)

	var (
		p = processor.NewFactory(processor.SVT_AV1, logger)
		s = scheduler.NewRoundRobin(1, p, logger)
	)

	for event := range eventChan {
		logger.Info(event.Event().String())

		if event.Event() != notify.Write && event.Event() != notify.Rename {
			continue
		}

		filename := filepath.Base(event.Path())

		if strings.HasPrefix(filename, config.Instance().ReleasePrefix) &&
			!strings.HasPrefix(filename, ".") {
			s.Schedule(ctx, scheduler.ConversionJob{
				InputFile: event.Path(),
			})
		}
	}
}
