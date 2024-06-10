package serverHttp

import (
	"log"
	"modEventFeedback/internal/controler/serverHttp"
	"net/http"
)

type Server struct {
	server http.Server
	host   string
	port   string
}

func loadRouters(router *http.ServeMux) {
	h := handler.New()

	router.HandleFunc("Get /", h.Get)
}

func New(host, post string) *Server {
	router := http.NewServeMux()
	loadRouters(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	return &Server{
		server: server,
		host:   host,
		port:   post,
	}
}

func Start(s *Server) {
	log.Fatalf("Server started. Lissening on port %s", s.port)
	s.server.ListenAndServe()
}
