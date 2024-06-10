package handler

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func New(r *http.ServeMux) *Handler {
	newHandler := Handler{}
	r.HandleFunc("GET /", newHandler.Get)

	return &newHandler
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("work handler")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Hello, world!"))
}
