package service

import "context"

type User struct {
	Id       int64
	Name     string
	Password string
	Email    string
}

type RegisterRequest struct {
	User User
}

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Token string
}

type ValidateTokenRequest struct {
	Token string
}

type ValidateTokenResponse struct {
	IsValid bool
}

type AuthService interface {
	Register(ctx context.Context, request *RegisterRequest) error
	Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error)
	ValidateToken(ctx context.Context, request *ValidateTokenRequest) (*ValidateTokenResponse, error)
}
