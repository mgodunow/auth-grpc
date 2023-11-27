package auth

import (
	"context"

	ssov1 "github.com/mgodunow/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	Login(ctx context.Context, email, password string, appId int) (string, error)
	Register(ctx context.Context, email, password string) (int64, error)
	IsAdmin(ctx context.Context, userId int64) (bool, error)
}

type serverAPI struct {
	ssov1.UnimplementedAuthServer
	auth Auth
}

func RegisterServerAPI(gRPC *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{auth: auth})
}

const (
	emptyValue  = 0
	emptyString = ""
)

func (s *serverAPI) Login(ctx context.Context, request *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	if err := validateLogin(request); err != nil {
		return nil, err
	}
	token, err := s.auth.Login(ctx, request.GetEmail(), request.GetPassword(), int(request.GetAppId()))
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &ssov1.LoginResponse{Token: token,}, nil
}

func validateLogin(request *ssov1.LoginRequest) error {
	if request.GetEmail() == emptyString {
		return status.Error(codes.InvalidArgument, "email id is required")
	}

	if request.GetPassword() == emptyString {
		return status.Error(codes.InvalidArgument, "password id is required")
	}

	if request.GetAppId() == emptyValue {
		return status.Error(codes.InvalidArgument, "app id is required")
	}
	return nil
}

func (s *serverAPI) Register(ctx context.Context, request *ssov1.RegisterRequest) (*ssov1.RegisterResponse, error) {
	if err := validateRegister(request); err != nil {
		return nil, err
	}

	userId, err := s.auth.Register(ctx, request.GetEmail(), request.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &ssov1.RegisterResponse{UserId: userId}, nil
}

func validateRegister(request *ssov1.RegisterRequest) error {
	if request.GetEmail() == emptyString {
		return status.Error(codes.InvalidArgument, "email id is required")
	}

	if request.GetPassword() == emptyString {
		return status.Error(codes.InvalidArgument, "password id is required")
	}

	return nil
}

func (s *serverAPI) IsAdmin(ctx context.Context, request *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	if err := validateIsAdmin(request); err != nil {
		return nil, err
	}
	isAdmin, err := s.auth.IsAdmin(ctx, request.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &ssov1.IsAdminResponse{IsAdmin: isAdmin}, nil
}

func validateIsAdmin(request *ssov1.IsAdminRequest) error {
	if request.GetUserId() == emptyValue {
		return status.Error(codes.InvalidArgument, "user_id is required")
	}
	return nil
}
