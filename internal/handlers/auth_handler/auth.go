package auth_handler

import (
	"context"
	"fmt"
	_ "fmt"
	"github.com/arrowwhi/message-auth/internal/interfaces/service"
	_ "google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/arrowwhi/message-auth/proto/auth/v1"
)

func (s *Service) Register(ctx context.Context, request *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	err := s.service.Register(ctx, &service.RegisterRequest{User: *s.cvt.RegisterToService(request)})
	if err != nil {
		return nil, fmt.Errorf("register to service: %w", err)
	}
	return &pb.SignUpResponse{Message: "ok"}, nil
}

func (s *Service) Login(ctx context.Context, request *pb.SignInRequest) (*pb.SignInResponse, error) {
	resp, err := s.service.Login(ctx, s.cvt.LoginToService(request))
	if err != nil {
		return nil, fmt.Errorf("login to service: %w", err)
	}
	return s.cvt.LoginToHandler(resp), nil
}

func (s *Service) ValidateToken(ctx context.Context, request *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	resp, err := s.service.ValidateToken(ctx, s.cvt.ValidateToService(request))
	if err != nil {
		return nil, fmt.Errorf("validate token to service: %w", err)
	}
	return s.cvt.ValidateToHandler(resp), nil
}
