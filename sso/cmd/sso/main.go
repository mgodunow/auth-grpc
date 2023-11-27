package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/mgodunow/auth-grpc/internal/app"
	"github.com/mgodunow/auth-grpc/internal/config"
)

func main() {
	config := config.MustLoad()

	log := setupLogger(config.Env)
	log.Info("Starting app")

	application := app.New(log, config.GRPC.Port, config.StoragePath, config.TokenTTL)
	go application.GRPCServer.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()
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
