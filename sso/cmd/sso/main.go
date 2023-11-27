package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/mgodunow/auth-grpc/internal/config"
)

func main() {
	config := config.MustLoad()
	fmt.Println(*config)
	//TODO: apps init
	log := setupLogger(config.Env)
	log.Info("Starting app")

	//TODO: run grpc-server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "development":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "production":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}
