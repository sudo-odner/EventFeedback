package serverHttp

import (
	"log/slog"
	"modEventFeedback/internal/config"
	"modEventFeedback/internal/controler/serverHttp/handler"
	"net/http"
)

type Server struct {
	server *http.Server
	cfg    *config.HTTPServer
	log    *slog.Logger
}

func NewServer(cfg *config.HTTPServer, log *slog.Logger) *Server {
	router := http.NewServeMux()
	handler.New(router)

	server := http.Server{
		Addr:    cfg.Host + ":" + cfg.Port,
		Handler: router,
	}
	return &Server{
		server: &server,
		cfg:    cfg,
		log:    log,
	}
}

func (s *Server) Start() {
	s.log.Info("Server started. Lissening on host " + s.cfg.Host + " port " + s.cfg.Port)

	s.server.ListenAndServe()
}
