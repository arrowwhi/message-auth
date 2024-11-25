package ep

import (
	"context"
	"fmt"
	"github.com/arrowwhi/go-utils/grpcserver"
	"github.com/arrowwhi/message-auth/internal/config"
	"github.com/arrowwhi/message-auth/internal/converter/converter"
	"github.com/arrowwhi/message-auth/internal/handlers/auth_handler"
	"github.com/arrowwhi/message-auth/internal/infra/postgres/auth_repo"
	"github.com/arrowwhi/message-auth/internal/service/auth_service_impl"
	"go.uber.org/zap"
)

func Run(cfg *config.Config, logger *zap.Logger) error {
	ctx := context.Background()

	pg, err := auth_repo.New(logger, cfg.Postgres)
	if err != nil {
		return err
	}

	authService := auth_service_impl.New(logger, pg, cfg.SecretKey, &converter.RepoConverterImpl{})

	srv, err := grpcserver.NewServer(
		cfg.Config,
		logger,
		grpcserver.WithImplementationAdapters(
			auth_handler.New(logger, authService, &converter.ServiceConverterImpl{}),
		),
		grpcserver.WithGrpcUnaryServerInterceptors(),
	)
	if err != nil {
		return fmt.Errorf("create gRPC server: %w", err)
	}

	logger.Info("Starting gRPC server", zap.String("gRPC port", cfg.Config.GRPCPort))
	err = srv.Start(ctx)
	if err != nil {
		logger.Error("Failed to start gRPC server", zap.Error(err))
	}

	return nil
}
