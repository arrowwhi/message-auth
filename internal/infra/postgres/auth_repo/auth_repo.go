package auth_repo

import (
	"context"
	"fmt"
	pg_lib "github.com/arrowwhi/go-utils/postgres"
	pg_config "github.com/arrowwhi/go-utils/postgres/db_config"
	"github.com/arrowwhi/message-auth/internal/interfaces/infra/postgres"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

var _ postgres.DatabaseRepo = (*impl)(nil)

type impl struct {
	logger *zap.Logger
	db     *pg_lib.Database
}

func New(logger *zap.Logger, config pg_config.DBConfig) (postgres.DatabaseRepo, error) {
	logger.Info("Starting database")
	m, err := pg_lib.NewDatabase(config)
	if err != nil {
		return nil, fmt.Errorf("postgres repo init: %w", err)
	}
	return &impl{
		logger: logger,
		db:     m,
	}, nil
}

func (i *impl) Get(ctx context.Context, email string) (*postgres.User, error) {
	query := `select id, name, email, password from users where email=$1;`
	args := []interface{}{email}

	rows, err := i.db.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	users := make([]*postgres.User, 0)

	for rows.Next() {
		var elem postgres.User
		if err = rows.Scan(
			&elem.Id, &elem.Name,
			&elem.Email, &elem.Password,
		); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}
		users = append(users, &elem)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows: %w", err)
	}
	if len(users) == 0 {
		return nil, pgx.ErrNoRows
	}
	return users[0], nil

}

func (i *impl) Add(ctx context.Context, user *postgres.User) error {
	//_ = i.db
	queryBuilder := i.db.Builder.Insert("users").
		Columns("name", "email", "password").
		Values(user.Name, user.Email, user.Password)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("queryBuilder.ToSql: %w", err)
	}
	i.logger.Debug(query)

	_, err = i.db.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("Pool.QueryRow: %w", err)
	}
	return nil
}
