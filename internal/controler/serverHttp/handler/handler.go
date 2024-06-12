package handler

import (
	"log/slog"
	"net/http"
)

type Handler struct {
	log *slog.Logger
}

func New(log *slog.Logger) *http.ServeMux {
	newHandler := Handler{
		log: log,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", newHandler.Get)
	mux.HandleFunc("GET /v1.0/health", newHandler.Health)
	mux.HandleFunc("GET /v1.0/some", newHandler.Health)

	return mux
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Hello"))
}
