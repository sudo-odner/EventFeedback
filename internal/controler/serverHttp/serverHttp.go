package serverHttp

import (
	"fmt"
	"net/http"
)

type Server struct {
	server http.Server
	host   string
	port   string
}

func loadRouters(router *http.ServeMux) {
	h := NewHandler()
	router.HandleFunc("GET /", h.Get)
}

func NewServer(host, port string) *Server {
	router := http.NewServeMux()
	loadRouters(router)

	server := http.Server{
		Addr:    host + ":" + port,
		Handler: router,
	}
	return &Server{
		server: server,
		host:   host,
		port:   port,
	}
}

func Start(s *Server) {
	fmt.Println("Server started. Lissening on port %s", s.port)
	s.server.ListenAndServe()
}
