package service

import (
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/anaxaim/tui/server/pkg/model"
	"github.com/anaxaim/tui/server/pkg/repository"
)

type moduleService struct {
	moduleRepository repository.ModuleRepository
}

func NewModuleService(moduleRepository repository.ModuleRepository) ModuleService {
	return &moduleService{
		moduleRepository: moduleRepository,
	}
}

func (m *moduleService) List() (model.TerraformModules, error) {
	return m.moduleRepository.List()
}

func (m *moduleService) Create(module *model.TerraformModule) (*model.TerraformModule, error) {
	parsedURL, err := url.Parse(module.GitRepositoryURL)
	if err != nil {
		return nil, err
	}

	domain := parsedURL.Hostname()
	path := strings.TrimLeft(parsedURL.Path, "/")
	if strings.Contains(domain, string(model.GITHUB)) {
		module.RegistryDetails.RegistryType = model.GITHUB
		module.RegistryDetails.ProjectID = path
	} else if strings.Contains(domain, string(model.GITLAB)) {
		module.RegistryDetails.RegistryType = model.GITLAB
		module.RegistryDetails.ProjectID = path
	}

	currentTime := time.Now()
	module.CreatedAt = currentTime
	module.CreatedAtString = currentTime.Format("15:04:05 02/01/2006")
	module.Status = model.CREATED

	return m.moduleRepository.Create(module)
}

func (m *moduleService) Get(id string) (*model.TerraformModule, error) {
	return m.GetModuleByID(id)
}

func (m *moduleService) Update(id string, newModule *model.TerraformModule) (*model.TerraformModule, error) {
	old, err := m.GetModuleByID(id)
	if err != nil {
		return nil, err
	}

	newModule.ID = old.ID
	currentTime := time.Now()
	newModule.UpdatedAt = &currentTime
	newModule.UpdatedAtString = currentTime.Format("15:04:05 02/01/2006")
	newModule.Status = model.UPDATED

	return m.moduleRepository.Update(newModule)
}

func (m *moduleService) Delete(id string) error {
	module, err := m.GetModuleByID(id)
	if err != nil {
		return err
	}

	return m.moduleRepository.Delete(module)
}

func (m *moduleService) Validate(module *model.TerraformModule) error {
	if module == nil {
		return errors.New("user is empty")
	}

	return nil
}

func (m *moduleService) GetModuleByID(id string) (*model.TerraformModule, error) {
	return m.moduleRepository.GetModuleByID(id)
}
