package repository

import (
	"context"

	"github.com/anaxaim/tui/server/pkg/model"
)

type Repository interface {
	User() UserRepository
	Module() ModuleRepository
	Credential() CredentialRepository
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

type CredentialRepository interface {
	GetCredentialByID(string) (*model.Credential, error)
	GetCredentialByName(string) (*model.Credential, error)
	List() (model.Credentials, error)
	Create(*model.Credential) (*model.Credential, error)
	Delete(*model.Credential) error
	Migrate() error
}
