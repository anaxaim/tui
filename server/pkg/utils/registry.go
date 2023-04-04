package utils

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"

	"github.com/anaxaim/tui/server/pkg/model"
)

func CloneGitRepo(module *model.TerraformModule) (*git.Repository, error) {
	memStorage := memory.NewStorage()
	repo, err := git.Clone(memStorage, nil, &git.CloneOptions{
		URL:           module.GitRepositoryURL,
		ReferenceName: plumbing.NewBranchReferenceName("master"),
		SingleBranch:  true,
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

func GetModuleContent(tree *object.Tree) (map[string]string, error) {
	content := make(map[string]string)
	err := tree.Files().ForEach(func(f *object.File) error {
		if !strings.HasPrefix(f.Name, ".terraform") && filepath.Ext(f.Name) == ".tf" {
			reader, err := f.Reader()
			if err != nil {
				return err
			}

			bytes, err := io.ReadAll(reader)
			if err != nil {
				return err
			}

			content[f.Name] = string(bytes)
		}

		return nil
	})

	return content, err
}

func ExtractRegistryFilesToDir(files map[string]string, destDir string) error {
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return err
	}

	for fileName, fileContent := range files {
		destPath := filepath.Join(destDir, fileName)
		if err := os.WriteFile(destPath, []byte(fileContent), 0o600); err != nil {
			return err
		}
	}

	return nil
}
