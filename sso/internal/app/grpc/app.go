package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	authgrpc "github.com/mgodunow/auth-grpc/internal/grpc/auth"
	"google.golang.org/grpc"
)

type App struct {
	log         *slog.Logger
	gRRPCServer *grpc.Server
	port        string
}

func New(log *slog.Logger, port string, authService authgrpc.Auth) *App {
	gRPCServer := grpc.NewServer()

	authgrpc.RegisterServerAPI(gRPCServer, authService)

	return &App{
		log:         log,
		gRRPCServer: gRPCServer,
		port:        port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"
	log := a.log.With(slog.String("op", op), slog.String("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%s",a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("grpc server is running", slog.String("address", l.Addr().String()))
	if err = a.gRRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).Info("Stopping gRPC server", slog.String("port", a.port))

	a.gRRPCServer.GracefulStop()
}
