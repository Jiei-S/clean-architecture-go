package usecase

import (
	"context"

	"github.com/Jiei-S/boilerplate-clean-architecture/internal/domain/entity"
	pkgErr "github.com/Jiei-S/boilerplate-clean-architecture/pkg/error"
)

type UserRepository interface {
	Save(ctx context.Context, e *entity.User) (*entity.User, *pkgErr.ApplicationError)
	Find(ctx context.Context, id string) (*entity.User, *pkgErr.ApplicationError)
}
