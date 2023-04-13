package utils

import (
	"fmt"
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
	branch := module.GitBranch

	if branch == "" {
		branch = "master"
	}

	referenceName := plumbing.NewBranchReferenceName(branch)

	repo, err := git.Clone(memStorage, nil, &git.CloneOptions{
		URL:           module.GitRepositoryURL,
		ReferenceName: referenceName,
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

func processFiles(tree *object.Tree, filterFunc func(string) bool) (map[string]string, error) {
	content := make(map[string]string)
	err := tree.Files().ForEach(func(f *object.File) error {
		if filterFunc(f.Name) {
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

func GetModuleContentRoot(tree *object.Tree) (map[string]string, error) {
	return processFiles(tree, func(filename string) bool {
		return !strings.HasPrefix(filename, ".terraform") && filepath.Ext(filename) == ".tf"
	})
}

func GetModuleContentDirectory(tree *object.Tree, moduleDirectory string) (map[string]string, error) {
	content, err := processFiles(tree, func(filename string) bool {
		if !strings.HasPrefix(filename, moduleDirectory+"/") {
			return false
		}

		relativePath := strings.TrimPrefix(filename, moduleDirectory+"/")

		return !strings.HasPrefix(relativePath, ".terraform") && filepath.Ext(relativePath) == ".tf"
	})
	if err != nil {
		return nil, err
	}

	trimmedContent := make(map[string]string)

	for key, value := range content {
		trimmedKey := strings.TrimPrefix(key, moduleDirectory+"/")
		trimmedContent[trimmedKey] = value
	}

	return trimmedContent, nil
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
