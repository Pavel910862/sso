package auth

import (
	"context"

	"github.com/Pavel910862/protos/protos/gen/go/sso"
	//sso "github.com/Pavel910862/protos/protos/gen/go/sso"
	"google.golang.org/grpc"
)

type serverAPI struct {
	sso.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
	sso.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Login(
	ctx context.Context,
	req *sso.LoginRequest,
) (*sso.LoginResponse, error) {
	return &sso.LoginResponse{
		Token: req.GetEmail(),
	}, nil
}

func (s *serverAPI) Register(
	ctx context.Context,
	req *sso.RegisterRequest,
) (*sso.RegisterResponse, error) {
	panic("implement me")
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	req *sso.IsAdminRequest,
) (*sso.IsAdminResponse, error) {
	panic("implement me")
}
