package converter

import (
	"github.com/arrowwhi/message-auth/internal/interfaces/infra/postgres"
	"github.com/arrowwhi/message-auth/internal/interfaces/service"
)

// goverter:converter
// goverter:output:file ./converter/converter.gen.go
// goverter:output:package :converter
// goverter:useZeroValueOnPointerInconsistency
// goverter:ignoreUnexported
// goverter:matchIgnoreCase
//
//go:generate goverter gen ./
type RepoConverter interface {
	UserToDatabase(user *service.User) *postgres.User
	UserToService(user *postgres.User) *service.User
}
