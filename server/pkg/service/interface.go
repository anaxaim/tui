package service

import (
	"github.com/anaxaim/tui/server/pkg/model"
)

type UserService interface {
	List() (model.Users, error)
	Create(*model.User) (*model.User, error)
	Get(string) (*model.User, error)
	Update(string, *model.User) (*model.User, error)
	Delete(string) error
	Validate(*model.User) error
	Auth(*model.AuthUser) (*model.User, error)
}