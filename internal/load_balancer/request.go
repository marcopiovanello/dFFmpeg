package loadbalancer

import "github.com/marcopeocchi/sanji/internal/ffmpeg/pb"

type Request struct {
	fn func() pb.FFmpeg_GetProgressClient
}

func NewRequest(fn func() pb.FFmpeg_GetProgressClient) Request {
	return Request{
		fn: fn,
	}
}
