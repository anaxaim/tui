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

type ModuleService interface {
	List() (model.TerraformModules, error)
	Get(string) (*model.TerraformModule, error)
	Create(*model.TerraformModule) (*model.TerraformModule, error)
	Update(string, *model.TerraformModule) (*model.TerraformModule, error)
	Delete(id string) error
	Validate(*model.TerraformModule) error
}
