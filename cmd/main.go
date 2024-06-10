package main

import (
	"log/slog"
	"modEventFeedback/internal/config"
	"modEventFeedback/internal/controler/serverHttp"
	"os"
)

func main() {
	// Иницилизация конфига при помощи пакета cleanenv
	cfg := config.MustLoad()

	// Иницилизация логера при помощи встроенного пакета slog
	log := setupLogger(cfg.Env)
	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug massages are enable")

	// TODO: init storage: mongoDB

	// TODO: init server: net/http

	// Иницилизация сервера, на основе HTTP протокола, и его ручек при помощи встроеного пакета http
	server := serverHttp.NewServer(cfg.HTTPServer.Host, cfg.HTTPServer.Port)
	serverHttp.Start(log, server)
}

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

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
