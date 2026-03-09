package main

import (
	"log/slog"
	"os"

	"github.com/Andrew1996-la/url-shortenerr/internal/config"
	"github.com/Andrew1996-la/url-shortenerr/internal/lib/logger/sl"
	"github.com/Andrew1996-la/url-shortenerr/internal/storage/sqlite"
)

const (
	localEnv = "local"
	devEnv   = "dev"
	prodEnv  = "prod"
)

func main() {
	cfg := config.MustLoad()
	logger := setupLogger(cfg.Env)
	storage, err := sqlite.NewStorage(cfg.StoragePath)
	if err != nil {
		logger.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	logger.Info("starting server", slog.String("env", cfg.Env))
}

// Конфигурация логгера пода разные окружения
func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case localEnv:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case devEnv:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case prodEnv:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
