package auth_handler

import (
	"context"
	"github.com/arrowwhi/message-auth/internal/converter"
	"github.com/arrowwhi/message-auth/internal/interfaces/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	pb "github.com/arrowwhi/message-auth/proto/auth/v1"
)

type Service struct {
	service service.AuthService
	logger  *zap.Logger
	cvt     converter.ServiceConverter
}

func New(logger *zap.Logger, service service.AuthService, cvt converter.ServiceConverter) *Service {
	return &Service{
		logger:  logger,
		service: service,
		cvt:     cvt,
	}
}

func (s *Service) RegisterServer(server *grpc.Server) {
	pb.RegisterAuthServiceServer(server, s)
}

func (s *Service) RegisterHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return pb.RegisterAuthServiceHandler(ctx, mux, conn)
}
