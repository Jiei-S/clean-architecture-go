package usecase

import (
	"context"

	"github.com/Jiei-S/boilerplate-clean-architecture/go-rest/internal/domain/entity"
	pkgErr "github.com/Jiei-S/boilerplate-clean-architecture/go-rest/pkg/error"
)

type UserRepository interface {
	Save(ctx context.Context, e *entity.User) (*entity.User, *pkgErr.ApplicationError)
	Find(ctx context.Context, id string) (*entity.User, *pkgErr.ApplicationError)
}
