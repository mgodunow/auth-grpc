package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/mgodunow/auth-grpc/internal/app/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, grpcPort string, storagePath string, tokenTTL time.Duration) *App {

	//TODO: storage init

	//TODO: init auth service

	grpcApp := grpcapp.New(log, grpcPort)
	return &App{
		GRPCServer: grpcApp,
	}
}
