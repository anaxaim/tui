package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type TerraformModule struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	GitRepositoryURL string             `json:"gitRepositoryUrl" bson:"gitRepositoryUrl"`
	GitBranch        string             `json:"gitBranch,omitempty" bson:"gitBranch,omitempty"`
	Directory        string             `json:"directory,omitempty" bson:"directory,omitempty"`
	Variables        []Variable         `json:"variables,omitempty" bson:"variables,omitempty"`
	Outputs          []Output           `json:"outputs,omitempty" bson:"outputs,omitempty"`
	Name             string             `json:"name" bson:"name"`
	Description      string             `json:"description,omitempty" bson:"description,omitempty"`
	TerraformImage   TerraformImage     `json:"terraformImage,omitempty" bson:"terraformImage,omitempty" `
	RegistryDetails  RegistryDetails    `json:"registryDetails,omitempty" bson:"registryDetails,omitempty"`
	MainProvider     string             `json:"mainProvider,omitempty" bson:"mainProvider,omitempty"`

	BaseModel
}

type Variable struct {
	Name            string `json:"name" bson:"name"`
	Type            string `json:"type,omitempty" bson:"type,omitempty"`
	Description     string `json:"description,omitempty" bson:"description,omitempty"`
	DefaultValue    string `json:"defaultValue,omitempty" bson:"defaultValue,omitempty"`
	Editable        bool   `json:"editable,omitempty" bson:"editable,omitempty"`
	Mandatory       bool   `json:"mandatory,omitempty" bson:"mandatory,omitempty"`
	ValidationRegex string `json:"validationRegex,omitempty" bson:"validationRegex,omitempty"`
}

type Output struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty" `
	Value       string `json:"value,omitempty" bson:"value,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Sensitive   string `json:"sensitive,omitempty" bson:"sensitive,omitempty"`
}

type TerraformImage struct {
	Repository string `json:"repository" bson:"repository"`
	Tag        string `json:"tag" bson:"tag"`
}

type RegistryDetails struct {
	RegistryType RegistryType `json:"registryType,omitempty" bson:"registryType,omitempty"`
	ProjectID    string       `json:"projectId,omitempty" bson:"projectId,omitempty"`
}

type TerraformModules []TerraformModule
