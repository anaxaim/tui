package service

import (
	"github.com/anaxaim/tui/backend/pkg/model"
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
	ImportModuleContent(id, workingDir string) (*model.TerraformModule, error)
	Execute(terraformVersion, command, id string) ([]byte, error)
	Validate(*model.TerraformModule) error
}

type CredentialService interface {
	List() (model.Credentials, error)
	Create(*model.Credential) (*model.Credential, error)
	Get(id string) (*model.Credential, error)
	Delete(id string) error
}
