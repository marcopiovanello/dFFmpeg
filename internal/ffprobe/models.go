package ffprobe

import (
	"strconv"
	"strings"
	"time"
)

type Tags struct {
	Language                 string `json:"language"`
	Title                    string `json:"title"`
	BPS                      string `json:"BPS"`
	DURATION                 string `json:"DURATION"`
	NUMBEROFFRAMES           string `json:"NUMBER_OF_FRAMES"`
	NUMBEROFBYTES            string `json:"NUMBER_OF_BYTES"`
	STATISTICSWRITINGAPP     string `json:"_STATISTICS_WRITING_APP"`
	STATISTICSWRITINGDATEUTC string `json:"_STATISTICS_WRITING_DATE_UTC"`
	STATISTICSTAGS           string `json:"_STATISTICS_TAGS"`
}

// Represent either a video, an audio or subtitle stream.
//
// In order to distinguish them use the CodecType field value (video, audio, subtitle)
type Stream struct {
	Index              int    `json:"index"`
	CodecName          string `json:"codec_name"`
	CodecLongName      string `json:"codec_long_name"`
	Profile            string `json:"profile,omitempty"`
	CodecType          string `json:"codec_type"`
	CodecTagString     string `json:"codec_tag_string"`
	CodecTag           string `json:"codec_tag"`
	Width              int    `json:"width,omitempty"`
	Height             int    `json:"height,omitempty"`
	CodedWidth         int    `json:"coded_width,omitempty"`
	CodedHeight        int    `json:"coded_height,omitempty"`
	ClosedCaptions     int    `json:"closed_captions,omitempty"`
	FilmGrain          int    `json:"film_grain,omitempty"`
	HasBFrames         int    `json:"has_b_frames,omitempty"`
	SampleAspectRatio  string `json:"sample_aspect_ratio,omitempty"`
	DisplayAspectRatio string `json:"display_aspect_ratio,omitempty"`
	PixFmt             string `json:"pix_fmt,omitempty"`
	Level              int    `json:"level,omitempty"`
	ColorRange         string `json:"color_range,omitempty"`
	ColorSpace         string `json:"color_space,omitempty"`
	ColorTransfer      string `json:"color_transfer,omitempty"`
	ColorPrimaries     string `json:"color_primaries,omitempty"`
	ChromaLocation     string `json:"chroma_location,omitempty"`
	Refs               int    `json:"refs,omitempty"`
	RFrameRate         string `json:"r_frame_rate"`
	AvgFrameRate       string `json:"avg_frame_rate"`
	TimeBase           string `json:"time_base"`
	StartPts           int    `json:"start_pts"`
	StartTime          string `json:"start_time"`
	ExtradataSize      int    `json:"extradata_size"`
	Tags               Tags   `json:"tags"`
	SampleFmt          string `json:"sample_fmt,omitempty"`
	SampleRate         string `json:"sample_rate,omitempty"`
	Channels           int    `json:"channels,omitempty"`
	ChannelLayout      string `json:"channel_layout,omitempty"`
	BitsPerSample      int    `json:"bits_per_sample,omitempty"`
	InitialPadding     int    `json:"initial_padding,omitempty"`
	BitsPerRawSample   string `json:"bits_per_raw_sample,omitempty"`
	DurationTs         int    `json:"duration_ts,omitempty"`
	Duration           string `json:"duration,omitempty"`
}

func (s *Stream) IsAudio() bool {
	return s.CodecType == "audio"
}

func (s *Stream) IsVideo() bool {
	return s.CodecType == "video"
}

func (s *Stream) IsSubtitle() bool {
	return s.CodecType == "subtitle"
}

func (s *Stream) ParseFPS() (int, error) {
	values := strings.Split(s.AvgFrameRate, "/")
	if len(values) < 2 && len(values) > 0 {
		return strconv.Atoi(s.AvgFrameRate)
	}

	num, err := strconv.Atoi(values[0])
	if err != nil {
		return 0, err
	}

	den, err := strconv.Atoi(values[1])
	if err != nil {
		return 0, err
	}

	return int(num / den), nil
}

func (s *Stream) ParseDuration() (time.Duration, error) {
	t, err := time.Parse("15:04:05", s.Tags.DURATION)
	if err != nil {
		return time.Duration(0), err
	}
	t = t.AddDate(1970, 0, 0)

	return time.Duration(t.UnixNano()), nil
}

func (s *Stream) TotalFrames() (int64, error) {
	fps, err := s.ParseFPS()
	if err != nil {
		return 0, err
	}

	duration, err := s.ParseDuration()
	if err != nil {
		duration = time.Duration(s.DurationTs * int(time.Millisecond))
	}

	return int64(fps * int(duration.Seconds())), nil
}

type Format struct {
	Filename       string `json:"filename"`
	NbStreams      int    `json:"nb_streams"`
	NbPrograms     int    `json:"nb_programs"`
	NbStreamGroups int    `json:"nb_stream_groups"`
	FormatName     string `json:"format_name"`
	FormatLongName string `json:"format_long_name"`
	StartTime      string `json:"start_time"`
	Duration       string `json:"duration"`
	Size           string `json:"size"`
	BitRate        string `json:"bit_rate"`
	ProbeScore     int    `json:"probe_score"`
}

func (f *Format) ParseDuration() (time.Duration, error) {
	fDuration, err := strconv.ParseFloat(f.Duration, 64)
	if err != nil {
		return time.Duration(0), err
	}

	return time.Duration(fDuration * float64(time.Second)), nil
}

type FFprobeOutput struct {
	Streams []Stream `json:"streams"`
	Format  Format   `json:"format"`
}

func (f *FFprobeOutput) GetVideoStream() *Stream {
	for _, stream := range f.Streams {
		if stream.IsVideo() {
			return &stream
		}
	}
	return &f.Streams[0]
}

func (f *FFprobeOutput) GetSubtitleStream() *Stream {
	for _, stream := range f.Streams {
		if stream.IsSubtitle() {
			return &stream
		}
	}
	return &f.Streams[0]
}
