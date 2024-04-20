package processor

import (
	"log/slog"

	"github.com/marcopeocchi/sanji/config"
)

const (
	SVT_AV1 = iota
	HEVC_QSV
)

func NewFactory(profile int, logger *slog.Logger) VideoProcessor {
	ffmpeg := config.Instance().FFMpegPath

	switch profile {
	case SVT_AV1:
		return NewAV1Processor(ffmpeg, "6", logger)
	case HEVC_QSV:
		return NewHEVCQSVProcessor(ffmpeg, 20)
	default:
		return nil
	}
}
