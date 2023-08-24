package gateway

import (
	"context"
	"database/sql"
	"time"

	"github.com/Jiei-S/boilerplate-clean-architecture/go-rest/internal/domain/entity"
	pkgErr "github.com/Jiei-S/boilerplate-clean-architecture/go-rest/pkg/error"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	ID        string    `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	FirstName string    `bun:"first_name,notnull"`
	LastName  string    `bun:"last_name,notnull"`
	Age       int32     `bun:"age,notnull"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp"`
}

var _ bun.BeforeAppendModelHook = (*User)(nil)

func (u *User) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		uuidObj, _ := uuid.NewUUID()
		u.ID = uuidObj.String()
		return nil
	}
	return nil
}

func (u *User) ToEntity() *entity.User {
	return &entity.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
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
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func RepositoryError(err error) *pkgErr.ApplicationError {
	switch err {
	case sql.ErrNoRows:
		return pkgErr.NewApplicationError(err.Error(), pkgErr.LevelWarn, pkgErr.CodeNotFound)
	default:
		return pkgErr.NewApplicationError(err.Error(), pkgErr.LevelError, pkgErr.CodeInternalServerError)
	}
}
