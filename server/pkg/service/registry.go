package service

import (
	"io"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"

	"github.com/anaxaim/tui/server/pkg/model"
	"github.com/anaxaim/tui/server/pkg/repository"
)

type registryService struct {
	registryRepository repository.RegistryRepository
	moduleRepository   repository.ModuleRepository
}

func NewRegistryService(registryRepository repository.RegistryRepository, moduleRepository repository.ModuleRepository) RegistryService {
	return &registryService{
		registryRepository: registryRepository,
		moduleRepository:   moduleRepository,
	}
}

func (r *registryService) ImportModuleContentByID(id string) (*model.RegistryContent, error) {
	module, err := r.moduleRepository.GetModuleByID(id)
	if err != nil {
		return nil, err
	}

	registry := new(model.RegistryContent)
	registry.RegistryType = module.RegistryDetails.RegistryType
	registry.ModuleID = module.ID

	// Clone git repository into memory
	memStorage := memory.NewStorage()
	repo, err := git.Clone(memStorage, nil, &git.CloneOptions{
		URL:           module.GitRepositoryURL,
		ReferenceName: plumbing.NewBranchReferenceName("master"),
		SingleBranch:  true,
	})
	if err != nil {
		return nil, err
	}

	ref, err := repo.Head()
	if err != nil {
		return nil, err
	}

	// Get the commit object
	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return nil, err
	}

	tree, err := commit.Tree()
	if err != nil {
		return nil, err
	}

	content := make(map[string]string)
	err = tree.Files().ForEach(func(f *object.File) error {
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

	if err != nil {
		return nil, err
	}

	registry.Content = content

	return r.registryRepository.Save(registry)
}

func (r *registryService) GetModuleContentByID(id string) (*model.RegistryContent, error) {
	return r.registryRepository.Get(id)
}
