package model

const GITHUB = "github"
const GITLAB = "gitlab"

type RegistryType struct {
	repositoriesUrl string
	repositoryUrl   string
	fileContentUrl  string
	readmeUrl       string
}

var GitHubRegistryType = RegistryType{
	repositoriesUrl: "https://api.github.com/user/repos?visibility=public&per_page=100",
	repositoryUrl:   "https://api.github.com/repos/%s",
}

var GitLabRegistryType = RegistryType{
	repositoriesUrl: "https://gitlab.com/api/v4/projects?visibility=public&owned=true",
	repositoryUrl:   "https://gitlab.com/api/v4/projects/%s",
}

type RegistryDetails struct {
	RegistryType string `json:"registryType,omitempty" bson:"registryType,omitempty"`
	ProjectID    string `json:"projectId,omitempty" bson:"projectId,omitempty"`
}
