package converter

import (
	"github.com/arrowwhi/message-auth/internal/interfaces/service"
	v1 "github.com/arrowwhi/message-auth/proto/auth/v1"
)

// goverter:converter
// goverter:output:file ./converter/converter.gen.go
// goverter:output:package :converter
// goverter:useZeroValueOnPointerInconsistency
// goverter:ignoreUnexported
// goverter:matchIgnoreCase
//
//go:generate goverter gen ./
type ServiceConverter interface {
	// goverter:ignore Id
	RegisterToService(req *v1.SignUpRequest) *service.User
	LoginToService(req *v1.SignInRequest) *service.LoginRequest
	// goverter:map Token AccessToken
	LoginToHandler(resp *service.LoginResponse) *v1.SignInResponse
	// goverter:map AccessToken Token
	ValidateToService(req *v1.ValidateTokenRequest) *service.ValidateTokenRequest
	ValidateToHandler(resp *service.ValidateTokenResponse) *v1.ValidateTokenResponse
}
