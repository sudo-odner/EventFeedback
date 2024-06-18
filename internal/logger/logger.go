package logger

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type Logger struct {
	Slog *slog.Logger
}

func makeLog(w io.Writer, level slog.Level) *slog.Logger {
	return slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		AddSource: true,
		Level:     &level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				source, _ := a.Value.Any().(*slog.Source)
				if source != nil {
					source.File = filepath.Base(source.File)
				}
			}
			return a
		},
	}))
}

func New(env string) *Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = makeLog(os.Stderr, slog.LevelDebug)
	case envDev:
		log = makeLog(os.Stderr, slog.LevelDebug)
	case envProd:
		log = makeLog(os.Stderr, slog.LevelInfo)
	}
	log.Info("Logger start work", slog.String("env", env))
	log.Debug("Debug massage are enable")

	return &Logger{
		Slog: log,
	}
}
