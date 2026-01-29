package main

import (
	"log/slog"
	"os"
	"sso/internal/app"
	"sso/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	application.gRPCSrv.MustRun()

	// TODO: init app
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case dev:
		log = slog.New(
			slog.NewJsonHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case prod:
		log = slog.New(
			slog.NewJsonHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}
	return log
}
