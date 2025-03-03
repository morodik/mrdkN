package auth

import (
	"context"
	"log"

	ssov1 "gen/go/sso"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
}

func (s *serverAPI) Register(ctx context.Context, req *ssov1.RegistertRequest) (*ssov1.RegisterResponse, error) {
	log.Println("Recived register request:", req.Email)
	return &ssov1.RegisterResponse{
		UserId: 123,
	}, nil
}
func (s *serverAPI) Login(ctx context.Context, req *ssov1.LoginRequest) (*ssov1.LoginResponse, error) {
	log.Println("Recived login request:", req.Email)
	return &ssov1.LoginResponse{
		Token: "some_generated_token",
	}, nil
}
func (s *serverAPI) IsAdmin(ctx context.Context, req *ssov1.IsAdminRequest) (*ssov1.IsAdminResponse, error) {
	log.Println("Received IsAdmin request for user ID", req.UserId)
	return &ssov1.IsAdminResponse{
		IsAdmin: true,
	}, nil
}
