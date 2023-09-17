package usecase

import (
	"github.com/Jiei-S/boilerplate-clean-architecture/internal/domain/entity"
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	Age       int32
}

func (u *User) ToEntity() *entity.User {
	return &entity.User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
	}
}

func FromEntity(
	entity *entity.User,
) *User {
	return &User{
		ID:        entity.ID,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Age:       entity.Age,
	}
}
