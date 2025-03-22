package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/marcopeocchi/sanji/internal/ffmpeg/pb"
	"github.com/marcopeocchi/sanji/internal/processor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var filename string

func main() {
	flag.StringVar(&filename, "f", "", "file path to send")
	flag.Parse()

	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", 2008), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	log.Fatalln(testStreaming(conn))
}

func testStreaming(conn *grpc.ClientConn) error {
	client := pb.NewFFmpegClient(conn)

	id := uuid.NewString()

	res, err := client.PrepareConversion(context.Background(), &pb.PrepareConversionRequest{
		Id:                     id,
		Filename:               filename,
		Processor:              int32(processor.RAV1E_AV1),
		AdditionalFfmpegParams: "",
	})
	if err != nil {
		return err
	}

	log.Println(res.TempFilename)

	fd, err := os.Open(filename)
	if err != nil {
		return err
	}

	streamingClient, err := client.StartConversion(context.Background())
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		br := bufio.NewReader(fd)
		buf := make([]byte, int64(2048*1000))

		for {
			_, err := br.Read(buf)
			if err != nil {
				wg.Done()
				log.Println("succesfully sent file", err)
				break
			}

			streamingClient.Send(&pb.ConversionContent{
				Id:          id,
				FileContent: buf,
			})
		}

		buf = nil
	}()

	go func() {
		defer wg.Done()

		for {
			res, err := streamingClient.Recv()
			if err == io.EOF {
				log.Fatalf("EOF")
				return
			}
			if len(res.FfmpegOutput) == 0 {
				log.Println("conversion ended")
				return
			}
			if err != nil {
				log.Fatalf("failed to receive: %v", err)
			}

			log.Println(string(res.FfmpegOutput))
		}
	}()

	go func() {
		log.Println("waiting for progress...")
		time.Sleep(3 * time.Second)

		progressClient, err := client.GetProgress(context.Background(), &pb.Query{
			Id: id,
		})
		if err != nil {
			log.Println(err)
		}

		go func() {
			for {
				res, err := progressClient.Recv()
				if err == io.EOF {
					log.Fatalf("EOF")
					return
				}
				if err != nil {
					log.Fatalf("failed to receive: %v", err)
				}

				log.Printf("percentage=%.2f fps=%d bitrate=%s",
					res.Ratio*100,
					res.Fps,
					res.BitRate,
				)
			}
		}()

		time.AfterFunc(5*time.Minute, func() {
			progressClient.CloseSend()
			wg.Done()
		})
	}()

	wg.Wait()
	streamingClient.CloseSend()
	log.Println("close send")

	return nil
}
