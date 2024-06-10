package serverHttp

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("work handler")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Hello, world!"))
}
