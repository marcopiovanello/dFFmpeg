package loadbalancer

import (
	"io"
)

type Worker struct {
	requests chan Request // conversions to do
	pending  int          // conversions pending
	index    int          // index in the heap
}

func (w *Worker) Work(done chan *Worker) {
	for {
		req := <-w.requests
		// TODO monitor grpc stream
		// signal done when stream ends
		for {
			_, err := req.fn().Recv()
			if err == io.EOF || err != nil {
				break
			}
		}
		done <- w
	}
}
