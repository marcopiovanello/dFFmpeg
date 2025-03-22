package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/marcopeocchi/sanji/internal/orchestrator"
)

type Handler struct {
	orc orchestrator.Orchestrator
}

func NewHandler(o orchestrator.Orchestrator) *Handler {
	return &Handler{
		orc: o,
	}
}

type StartJobRequest struct {
	Path    string `json:"path"`
	Encoder int    `json:"encoder"`
	Quality int    `json:"quality"`
}

func (h *Handler) StartJob() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json")

		req := StartJobRequest{}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id, err := h.orc.StartJob(context.Background(), req.Path, req.Encoder, req.Quality)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(id)
	}
}

func (h *Handler) Aggregate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "SSE not supported", http.StatusInternalServerError)
			return
		}

		progressChan, err := h.orc.Aggregate(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for {
			select {
			case <-r.Context().Done():
				log.Println("detaching from fan-in channel")
				return
			case msg, ok := <-progressChan:
				if !ok {
					http.Error(w, "closed fan-in channel", http.StatusInternalServerError)
					return
				}

				var b bytes.Buffer

				b.WriteString("event: progress\n")
				b.WriteString("data: ")

				if err := json.NewEncoder(&b).Encode(msg); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				b.WriteRune('\n')
				b.WriteRune('\n')

				io.Copy(w, &b)

				flusher.Flush()
			}
		}
	}
}

func (h *Handler) Details() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "SSE not supported", http.StatusInternalServerError)
			return
		}

		id := chi.URLParam(r, "id")

		progressChan, err := h.orc.Details(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for {
			select {
			case <-r.Context().Done():
				log.Println("detaching from fan-in channel")
				return
			case msg, ok := <-progressChan:
				if !ok {
					http.Error(w, "closed fan-in channel", http.StatusInternalServerError)
					return
				}

				var b bytes.Buffer

				b.WriteString("event: progress\n")
				b.WriteString("data: ")

				if err := json.NewEncoder(&b).Encode(msg); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				b.WriteRune('\n')
				b.WriteRune('\n')

				io.Copy(w, &b)

				flusher.Flush()
			}
		}
	}
}

func (h *Handler) StopJob() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json")

		id := chi.URLParam(r, "id")
		h.orc.StopJob(r.Context(), id)

		json.NewEncoder(w).Encode(id)
	}
}
