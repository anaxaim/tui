package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"

	"github.com/anaxaim/tui/server/pkg/model"
)

const GitlabAccessToken = "GITLAB_ACCESS_TOKEN"

func CloneGitRepo(module *model.TerraformModule, auth *http.BasicAuth) (*git.Repository, error) {
	memStorage := memory.NewStorage()
	branch := module.GitBranch

	if branch == "" {
		branch = "master"
	}

	referenceName := plumbing.NewBranchReferenceName(branch)

	repo, err := git.Clone(memStorage, nil, &git.CloneOptions{
		URL:           module.GitRepositoryURL,
		ReferenceName: referenceName,
		SingleBranch:  true,
		Auth:          auth,
	})

	return repo, err
}

func GetCommitTree(repo *git.Repository) (*object.Tree, error) {
	ref, err := repo.Head()
	if err != nil {
		return nil, err
	}

	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}

	tree, err := commit.Tree()

	return tree, err
}

func WriteFiles(tempDir string, content map[string]string) error {
	for fileName, fileContent := range content {
		filePath := filepath.Join(tempDir, fileName)
		if err := os.WriteFile(filePath, []byte(fileContent), 0o600); err != nil {
			return fmt.Errorf("failed to write file %s: %w", fileName, err)
		}
	}

	return nil
}

func SetGitlabAuth(auth *http.BasicAuth, secrets []model.Secret) {
	for _, v := range secrets {
		if v.Name == GitlabAccessToken {
			auth.Password = v.Value
			auth.Username = "tui.user"

			return
		}
	}
}
