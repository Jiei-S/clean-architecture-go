package usecase

import (
	"context"

	pkgErr "github.com/Jiei-S/boilerplate-clean-architecture/go-rest/pkg/error"
)

type UserUsecase interface {
	AddUser(ctx context.Context, dto *User) (*User, *pkgErr.ApplicationError)
	FindUser(ctx context.Context, id string) (*User, *pkgErr.ApplicationError)
}
