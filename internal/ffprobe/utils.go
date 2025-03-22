package ffprobe

import "time"

func TotalFrames(streams *[]Stream) (int64, error) {
	var (
		fps      int
		duration time.Duration
		err      error
	)

	for _, stream := range *streams {
		if stream.IsVideo() {
			fps, err = stream.ParseFPS()
			if err != nil {
				return 0, err
			}
			duration, _ = stream.ParseDuration()
		}
		if stream.IsSubtitle() && stream.DurationTs != 0 {
			duration = time.Duration(stream.DurationTs) * time.Millisecond
		}
	}

	return int64(fps * int(duration.Seconds())), nil
}
