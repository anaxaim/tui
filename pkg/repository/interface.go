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

type UserRepository interface {
	GetUserByUsername(string) (*model.User, error)
	List() (model.Users, error)
	Delete(*model.User) error
	Create(*model.User) (*model.User, error)
	Update(*model.User) (*model.User, error)
	Exists(string) (bool, error)
	Migrate() error
}
