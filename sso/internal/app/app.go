package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/mgodunow/auth-grpc/internal/app/grpc"
	"github.com/mgodunow/auth-grpc/internal/services/auth"
	"github.com/mgodunow/auth-grpc/internal/storage/sqlite"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort string, storagePath string, tokenTTL time.Duration) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, grpcPort, authService)
	return &App{
		GRPCServer: grpcApp,
	}
}
