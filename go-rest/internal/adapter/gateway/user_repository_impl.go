package gateway

import (
	"context"

	"github.com/Jiei-S/boilerplate-clean-architecture/go-rest/internal/domain/entity"
	"github.com/Jiei-S/boilerplate-clean-architecture/go-rest/internal/usecase"
	pkgErr "github.com/Jiei-S/boilerplate-clean-architecture/go-rest/pkg/error"
	"github.com/uptrace/bun"
)

type TxKey string

const TX_KEY TxKey = "xxxxx"

var _ usecase.UserRepository = (*UserRepositoryImpl)(nil)

type UserRepositoryImpl struct {
}

func (u *UserRepositoryImpl) Save(ctx context.Context, entity *entity.User) (*entity.User, *pkgErr.ApplicationError) {
	tx := ctx.Value(TX_KEY).(*bun.Tx)

	user := FromEntity(entity)
	if _, err := tx.NewInsert().Model(user).Exec(ctx); err != nil {
		return nil, RepositoryError(err)
	}
	return user.ToEntity(), nil
}

func (u *UserRepositoryImpl) Find(ctx context.Context, id string) (*entity.User, *pkgErr.ApplicationError) {
	tx := ctx.Value(TX_KEY).(*bun.Tx)

	var user User
	if err := tx.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx); err != nil {
		return nil, RepositoryError(err)
	}
	return user.ToEntity(), nil
}

func NewUserRepository() usecase.UserRepository {
	return &UserRepositoryImpl{}
}
