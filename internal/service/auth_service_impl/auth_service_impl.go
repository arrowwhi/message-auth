package auth_service_impl

import (
	"context"
	"errors"
	"fmt"
	"github.com/arrowwhi/message-auth/internal/converter"
	"github.com/arrowwhi/message-auth/internal/interfaces/infra/postgres"
	"github.com/arrowwhi/message-auth/internal/interfaces/service"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var _ service.AuthService = (*impl)(nil)

type impl struct {
	logger *zap.Logger

	database  postgres.DatabaseRepo
	secretKey string
	cvt       converter.RepoConverter
}

func New(
	logger *zap.Logger,
	database postgres.DatabaseRepo,
	secretKey string,
	cvt converter.RepoConverter,
) service.AuthService {
	return &impl{
		logger:    logger,
		database:  database,
		secretKey: secretKey,
		cvt:       cvt,
	}
}

func (i *impl) Register(ctx context.Context, request *service.RegisterRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hashing password: %w", err)
	}
	request.User.Password = string(hashedPassword)

	// Convert and save user to the database
	err = i.database.Add(ctx, i.cvt.UserToDatabase(&request.User))
	if err != nil {
		return fmt.Errorf("register user: %w", err)
	}
	return nil
}

func (i *impl) Login(ctx context.Context, request *service.LoginRequest) (*service.LoginResponse, error) {
	user, err := i.database.Get(ctx, request.Email)
	if err != nil {
		return nil, fmt.Errorf("retrieving user: %w", err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// Check if the password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := i.generateToken(user.Id)
	if err != nil {
		return nil, fmt.Errorf("generating token: %w", err)
	}

	return &service.LoginResponse{
		Token: token,
	}, nil

}

func (i *impl) ValidateToken(ctx context.Context, request *service.ValidateTokenRequest) (*service.ValidateTokenResponse, error) {
	// Parse and validate the token
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(request.Token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(i.secretKey), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Ensure token is not expired
	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	return &service.ValidateTokenResponse{
		IsValid: true,
		//UserID:  claims.Subject,
	}, nil
}

// generateToken creates a JWT token for a given user ID.
func (i *impl) generateToken(userID int64) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   fmt.Sprintf("%d", userID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token valid for 24 hours
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(i.secretKey))
}
