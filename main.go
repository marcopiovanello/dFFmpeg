package main

import (
	"context"

	"github.com/illarion/gonotify/v2"
	"github.com/marcopeocchi/sanji/config"
	"github.com/marcopeocchi/sanji/processor"
	"github.com/marcopeocchi/sanji/scheduler"
)

func main() {
	config.LoadFile("config.yaml")

	watcher, err := gonotify.NewDirWatcher(
		context.Background(),
		gonotify.IN_CLOSE_NOWRITE,
		config.Instance().Root,
	)
	if err != nil {
		panic(err)
	}

	ffmpegProcessor := processor.NewFFmpegProcessor("ffmpeg", "libsvtav1")

	s := scheduler.NewRoundRobinScheduler(1, ffmpegProcessor)

	for event := range watcher.C {
		s.Schedule(scheduler.ConversionJob{
			InputFile: event.Name,
		})
	}
}
