package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/marcopeocchi/sanji/internal/config"
	"github.com/marcopeocchi/sanji/internal/ffmpeg"
	"github.com/marcopeocchi/sanji/internal/ffmpeg/pb"
	"google.golang.org/grpc"
)

func main() {
	cfgPath := flag.String("c", "config.yml", "path of the config file")
	flag.Parse()

	config.LoadFile(*cfgPath)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 2008))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterFFmpegServer(grpcServer, ffmpeg.NewFFmpegServer())

	grpcServer.Serve(lis)
}
