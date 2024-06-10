package serverHttp

import (
	"log/slog"
	"modEventFeedback/internal/controler/serverHttp/handler"
	"net/http"
)

type Server struct {
	server http.Server
	host   string
	port   string
}

func NewServer(host, port string) *Server {
	router := http.NewServeMux()
	handler.New(router)

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

func Start(log *slog.Logger, s *Server) {
	log.Info("Server started. Lissening on host " + s.host + " port " + s.port)

	s.server.ListenAndServe()
}
