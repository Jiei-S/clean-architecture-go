package usecase

import (
	"context"

	pkgErr "github.com/Jiei-S/boilerplate-clean-architecture/go-rest/pkg/error"
)

var _ UserUsecase = (*UserUsecaseImpl)(nil)

type UserUsecaseImpl struct {
	userRepository UserRepository
}

func (u *UserUsecaseImpl) AddUser(
	ctx context.Context,
	dto *User,
) (*User, *pkgErr.ApplicationError) {
	entity, err := u.userRepository.Save(ctx, dto.ToEntity())
	if err != nil {
		return nil, err
	}
	return FromEntity(entity), nil
}

func (u *UserUsecaseImpl) FindUser(
	ctx context.Context,
	id string,
) (*User, *pkgErr.ApplicationError) {
	entity, err := u.userRepository.Find(ctx, id)
	if err != nil {
		return nil, err
	}
	return FromEntity(entity), nil
}

func NewUserUsecase(
	userRepository UserRepository,
) UserUsecase {
	return &UserUsecaseImpl{
		userRepository: userRepository,
	}
}
