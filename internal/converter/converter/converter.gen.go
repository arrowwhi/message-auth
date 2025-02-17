// Code generated by github.com/jmattheis/goverter, DO NOT EDIT.
//go:build !goverter

package converter

import (
	postgres "github.com/arrowwhi/message-auth/internal/interfaces/infra/postgres"
	service "github.com/arrowwhi/message-auth/internal/interfaces/service"
	v1 "github.com/arrowwhi/message-auth/proto/auth/v1"
)

type RepoConverterImpl struct{}

func (c *RepoConverterImpl) UserToDatabase(source *service.User) *postgres.User {
	var pPostgresUser *postgres.User
	if source != nil {
		var postgresUser postgres.User
		postgresUser.Id = (*source).Id
		postgresUser.Name = (*source).Name
		postgresUser.Email = (*source).Email
		postgresUser.Password = (*source).Password
		pPostgresUser = &postgresUser
	}
	return pPostgresUser
}
func (c *RepoConverterImpl) UserToService(source *postgres.User) *service.User {
	var pServiceUser *service.User
	if source != nil {
		var serviceUser service.User
		serviceUser.Id = (*source).Id
		serviceUser.Name = (*source).Name
		serviceUser.Password = (*source).Password
		serviceUser.Email = (*source).Email
		pServiceUser = &serviceUser
	}
	return pServiceUser
}

type ServiceConverterImpl struct{}

func (c *ServiceConverterImpl) LoginToHandler(source *service.LoginResponse) *v1.SignInResponse {
	var pAuthSignInResponse *v1.SignInResponse
	if source != nil {
		var authSignInResponse v1.SignInResponse
		authSignInResponse.AccessToken = (*source).Token
		pAuthSignInResponse = &authSignInResponse
	}
	return pAuthSignInResponse
}
func (c *ServiceConverterImpl) LoginToService(source *v1.SignInRequest) *service.LoginRequest {
	var pServiceLoginRequest *service.LoginRequest
	if source != nil {
		var serviceLoginRequest service.LoginRequest
		serviceLoginRequest.Email = (*source).Email
		serviceLoginRequest.Password = (*source).Password
		pServiceLoginRequest = &serviceLoginRequest
	}
	return pServiceLoginRequest
}
func (c *ServiceConverterImpl) RegisterToService(source *v1.SignUpRequest) *service.User {
	var pServiceUser *service.User
	if source != nil {
		var serviceUser service.User
		serviceUser.Name = (*source).Name
		serviceUser.Password = (*source).Password
		serviceUser.Email = (*source).Email
		pServiceUser = &serviceUser
	}
	return pServiceUser
}
func (c *ServiceConverterImpl) ValidateToHandler(source *service.ValidateTokenResponse) *v1.ValidateTokenResponse {
	var pAuthValidateTokenResponse *v1.ValidateTokenResponse
	if source != nil {
		var authValidateTokenResponse v1.ValidateTokenResponse
		authValidateTokenResponse.IsValid = (*source).IsValid
		pAuthValidateTokenResponse = &authValidateTokenResponse
	}
	return pAuthValidateTokenResponse
}
func (c *ServiceConverterImpl) ValidateToService(source *v1.ValidateTokenRequest) *service.ValidateTokenRequest {
	var pServiceValidateTokenRequest *service.ValidateTokenRequest
	if source != nil {
		var serviceValidateTokenRequest service.ValidateTokenRequest
		serviceValidateTokenRequest.Token = (*source).AccessToken
		pServiceValidateTokenRequest = &serviceValidateTokenRequest
	}
	return pServiceValidateTokenRequest
}
