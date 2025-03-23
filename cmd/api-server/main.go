package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/marcopeocchi/sanji/internal/ffmpeg/pb"
	"github.com/marcopeocchi/sanji/internal/orchestrator"
	"github.com/marcopeocchi/sanji/internal/rest"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	h2s := &http2.Server{}
	r := chi.NewRouter()

	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", 2008), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	orc, err := orchestrator.NewGrpcOrchestrator(pb.NewFFmpegClient(conn), 8)
	if err != nil {
		log.Fatalln(err)
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	r.Use(corsMiddleware.Handler)

	hand := rest.NewHandler(orc)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/d/{id}", hand.Details())
		r.Get("/aggregate", hand.Aggregate())
		r.Post("/start", hand.StartJob())
		r.Get("/stop/{id}", hand.StopJob())
		r.Get("/running", hand.GetNodes())
	})

	server := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: h2c.NewHandler(r, h2s),
	}

	server.ListenAndServe()
}
