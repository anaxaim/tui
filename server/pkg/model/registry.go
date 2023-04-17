package model

import "github.com/hashicorp/terraform-config-inspect/tfconfig"

type RegistryType string

const (
	GITHUB RegistryType = "github"
	GITLAB RegistryType = "gitlab"
)

type RegistryDetails struct {
	RegistryType  RegistryType      `json:"registryType,omitempty" bson:"registryType,omitempty"`
	ProjectID     string            `json:"projectId,omitempty" bson:"projectId,omitempty"`
	Credentials   string            `json:"credentials,omitempty" bson:"credentials,omitempty"`
	Content       map[string]string `json:"content" bson:"content"`
	ParsedContent *ParsedContent    `json:"parsedContent" bson:"parsedContent"`
}

type ParsedContent struct {
	Variables   map[string]*ParsedVariable               `json:"variables,omitempty" bson:"variables,omitempty"`
	Outputs     map[string]*ParsedOutput                 `json:"outputs,omitempty" bson:"outputs,omitempty"`
	Providers   map[string]*tfconfig.ProviderRequirement `json:"requiredProviders,omitempty" bson:"requiredProviders,omitempty"`
	Resources   map[string]*ParsedResource               `json:"resources,omitempty" bson:"resources,omitempty"`
	DataSources map[string]*ParsedResource               `json:"dataSources,omitempty" bson:"dataSources,omitempty"`
}

type ParsedResource struct {
	Provider string `json:"provider" bson:"provider"`
}

type ParsedOutput struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Sensitive   bool   `json:"sensitive,omitempty" bson:"sensitive,omitempty"`
}

type ParsedVariable struct {
	Name        string      `json:"name" bson:"name"`
	Type        string      `json:"type,omitempty" bson:"type,omitempty"`
	Description string      `json:"description,omitempty" bson:"description,omitempty"`
	Default     interface{} `json:"default" bson:"default"`
	Required    bool        `json:"required" bson:"required"`
	Sensitive   bool        `json:"sensitive,omitempty" bson:"sensitive,omitempty"`
}
