package processor

import "github.com/marcopeocchi/sanji/internal/config"

type Encoder = int

const (
	SVT_AV1 Encoder = iota
	RAV1E_AV1
	HEVC_QSV
	HEVC_VIDEOTOOLBOX
)

// TODO: parametrize quality
func NewFactory(profile Encoder, quality int) VideoProcessor {
	ffmpeg := config.Instance().FFmpegPath

	switch profile {
	case SVT_AV1:
		return NewAV1SVTProcessor(ffmpeg, quality) // 6
	case RAV1E_AV1:
		return NewRav1eAV1Processor(ffmpeg, quality) // 6
	case HEVC_QSV:
		return NewHEVCQSVProcessor(ffmpeg, quality) // 20
	case HEVC_VIDEOTOOLBOX:
		return NewHEVCVideoToolboxProcessor(ffmpeg, quality) // 65
	default:
		return nil
	}
}
