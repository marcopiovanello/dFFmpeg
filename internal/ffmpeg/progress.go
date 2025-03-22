package ffmpeg

import (
	"regexp"
	"strconv"
	"strings"
)

// example ffmpeg progress log
// 2025/03/21 09:44:22 frame=   26 fps= 13 q=22.0 size=       1KiB time=00:00:01.04 bitrate=   5.3kbits/s speed=0.518x

var (
	//multipleSpacesRegExp = regexp.MustCompile(`\s\s+`)
	//keyValuesRegExp    = regexp.MustCompile(`\w+=\s+\d+`)
	fpsValueRegExp     = regexp.MustCompile(`fps=(\s+|)\d+`)
	frameValueRegExp   = regexp.MustCompile(`frame=(\s+|)\d+`)
	bitrateValueRegExp = regexp.MustCompile(`bitrate=(\s+|)(\S+|\d+)`)
)

type Progress struct {
	BitRate string  `json:"bit_rate"`
	Ratio   float64 `json:"ratio"`
	FPS     int     `json:"fps"`
	Q       int     `json:"q"`
}

func getValue(s string, re *regexp.Regexp) string {
	match := strings.ReplaceAll(re.FindString(s), " ", "")
	parts := strings.Split(match, "=")

	if len(parts) > 1 {
		return parts[1]
	}
	return ""
}

func parseProgress(logEntry string, totalFrames int64) (*Progress, error) {
	var (
		fps     = getValue(logEntry, fpsValueRegExp)
		frame   = getValue(logEntry, frameValueRegExp)
		bitrate = getValue(logEntry, bitrateValueRegExp)
	)

	nFPS, err := strconv.Atoi(fps)
	if err != nil {
		return nil, err
	}
	nFrame, err := strconv.Atoi(frame)
	if err != nil {
		return nil, err
	}

	return &Progress{
		BitRate: bitrate,
		Ratio:   float64(nFrame) / float64(totalFrames),
		FPS:     nFPS,
		Q:       0,
	}, nil
}
