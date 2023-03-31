package repository

import (
	"context"

	"github.com/anaxaim/tui/server/pkg/model"
)

type Repository interface {
	User() UserRepository
	Module() ModuleRepository
	Registry() RegistryRepository
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

type ModuleRepository interface {
	GetModuleByID(string) (*model.TerraformModule, error)
	List() (model.TerraformModules, error)
	Create(*model.TerraformModule) (*model.TerraformModule, error)
	Update(*model.TerraformModule) (*model.TerraformModule, error)
	Delete(*model.TerraformModule) error
	Migrate() error
}

type RegistryRepository interface {
	Save(registry *model.RegistryContent) (*model.RegistryContent, error)
	Get(string) (*model.RegistryContent, error)
}
