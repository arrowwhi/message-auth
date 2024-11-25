package postgres

import "context"

type User struct {
	Id       int64
	Name     string
	Email    string
	Password string
}

type DatabaseRepo interface {
	Add(ctx context.Context, user *User) error
	Get(ctx context.Context, email string) (*User, error)
}
