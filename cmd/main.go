package main

import (
	"modEventFeedback/internal/config"
	"modEventFeedback/internal/controler/serverHttp"
	"modEventFeedback/internal/logger"
)

func main() {
	// Иницилизация конфига при помощи пакета cleanenv
	cfg, err := config.MustLoad()
	if err != nil {
		return
	}

	// Иницилизация логера при помощи встроенного пакета slog
	log := logger.New(cfg.Env)

	// Иницилизация mongoDB
	// db := mongoDB.New(&cfg.DB.MongoDB, log.Slog)

	// TODO: init server: net/http

	// Иницилизация сервера, на основе HTTP протокола, и его ручек при помощи встроеного пакета http
	server := serverHttp.NewServer(&cfg.HTTPServer, log.Slog)
	server.Start()
}
