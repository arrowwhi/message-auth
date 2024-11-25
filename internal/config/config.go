package config

import (
	"github.com/arrowwhi/go-utils/grpcserver/grpc_config"
	"github.com/arrowwhi/go-utils/postgres/db_config"
)

type Config struct {
	Config    grpc_config.Config
	LogLevel  string             `envconfig:"LOG_LEVEL" default:"debug"` // Уровень логирования
	Postgres  db_config.DBConfig `envconfig:"POSTGRES"`
	SecretKey string             `envconfig:"JWT_SECRET_KEY" required:"true"`
}
