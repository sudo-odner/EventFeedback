package main

import (
	"fmt"
	"log/slog"
	"modEventFeedback/internal/config"
	"modEventFeedback/internal/controler/serverHttp"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// TODO: init config: cleanenv
	cfg := config.MustLoad()
	fmt.Println(cfg)

	//TODO: init logger: slog

	log := setupLogger(cfg.Env)
	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug massages are enable")

	// TODO: init storage: mongoDB

	// TODO: init server: net/http

	// Init server on protocol HTTP
	server := serverHttp.NewServer(cfg.HTTPServer.Host, cfg.HTTPServer.Port)
	serverHttp.Start(server)

	// TODO: init router:

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log

}
