package repository

import (
	"context"

	"github.com/anaxaim/tui/pkg/model"
)

type Repository interface {
	User() UserRepository
	Close() error
	Ping(ctx context.Context) error
}

type UserRepository interface { //nolint: interfacebloat
	GetUserByID(string) (*model.User, error)
	GetUserByAuthID(authType, authID string) (*model.User, error)
	GetUserByName(string) (*model.User, error)
	List() (model.Users, error)
	Create(*model.User) (*model.User, error)
	Update(*model.User) (*model.User, error)
	Delete(*model.User) error
	AddAuthInfo(authInfo *model.AuthInfo) error
	DelAuthInfo(authInfo *model.AuthInfo) error
	Exists(string) (bool, error)
	Migrate() error
}
