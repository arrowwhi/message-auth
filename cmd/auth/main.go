package main

import (
	"github.com/arrowwhi/go-utils/logger"
	"github.com/arrowwhi/message-auth/internal/config"
	"github.com/arrowwhi/message-auth/internal/ep"
	"go.uber.org/zap"
	"log"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.GetConfigFromEnv()
	if err != nil {
		log.Fatalf("Failed to load configuration: %s\n", err.Error())
	}

	// Инициализация логгера
	zapLogger := logger.NewClientZapLogger(cfg.LogLevel, cfg.Config.ServiceName)

	// Запуск сервера
	if err = ep.Run(cfg, zapLogger); err != nil {
		zapLogger.Fatal("Run server failed: %s\n", zap.Error(err))
	}
}
