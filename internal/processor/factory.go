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
func NewFactory(profile Encoder, qp *QualityPreset) VideoProcessor {
	ffmpeg := config.Instance().FFmpegPath

	switch profile {
	case SVT_AV1:
		return NewAV1SVTProcessor(ffmpeg, qp) // 6
	case RAV1E_AV1:
		return NewRav1eAV1Processor(ffmpeg, qp) // 6
	case HEVC_QSV:
		return NewHEVCQSVProcessor(ffmpeg, qp) // 20
	case HEVC_VIDEOTOOLBOX:
		return NewHEVCVideoToolboxProcessor(ffmpeg, qp) // 65
	default:
		return nil
	}
}
