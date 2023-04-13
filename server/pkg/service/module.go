package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/url"
	"os/exec"
	"strings"
	"time"

	"github.com/mattn/go-colorable"

	"github.com/anaxaim/tui/server/pkg/model"
	"github.com/anaxaim/tui/server/pkg/repository"
	"github.com/anaxaim/tui/server/pkg/utils"
)

var ErrUserIsEmpty = errors.New("user is empty")

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

func (m *moduleService) ImportModuleContent(id, workingDir string) (*model.TerraformModule, error) { //nolint: cyclop
	module, err := m.GetModuleByID(id)
	if err != nil {
		return nil, err
	}

	repo, err := utils.CloneGitRepo(module)
	if err != nil {
		return nil, err
	}

	tree, err := utils.GetCommitTree(repo)
	if err != nil {
		return nil, err
	}

	var content map[string]string
	if module.Directory == "" {
		content, err = utils.GetModuleContentRoot(tree)
	} else {
		content, err = utils.GetModuleContentDirectory(tree, module.Directory)
	}

	if err != nil {
		return nil, err
	}

	if len(module.Variables) > 0 {
		contentVariables, ok := content["variables.tf"]
		if ok {
			newContentVariables, err := utils.UpdateContentVariables(contentVariables, module.Variables)
			if err != nil {
				return nil, err
			}

			content["variables.tf"] = newContentVariables
		}
	}

	if err := utils.WriteFiles(workingDir, content); err != nil {
		return nil, err
	}

	tfModule, err := utils.LoadTFModule(workingDir)
	if err != nil {
		return nil, err
	}

	module.RegistryDetails.Content = content
	module.RegistryDetails.ParsedContent = tfModule
	module.Status = model.RUNNING

	_, err = m.moduleRepository.Update(module)
	if err != nil {
		return nil, err
	}

	return module, nil
}

func (m *moduleService) Execute(terraformVersion, command, id string) ([]byte, error) {
	containerName := "tui-terraform-" + terraformVersion
	workingDir := fmt.Sprintf("/terraform/%s", id)

	cmd := exec.CommandContext(context.Background(), "docker", "exec", "-i", "-u", "0", "-w", workingDir, containerName, "terraform", command)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = colorable.NewNonColorable(&stdout)
	cmd.Stderr = colorable.NewNonColorable(&stderr)

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("terraform command execution failed: %w, stderr: %s", err, stderr.String())
	}

	return stdout.Bytes(), nil
}

func (m *moduleService) Validate(module *model.TerraformModule) error {
	if module == nil {
		return fmt.Errorf("%w", ErrUserIsEmpty)
	}

	return nil
}

func (m *moduleService) GetModuleByID(id string) (*model.TerraformModule, error) {
	return m.moduleRepository.GetModuleByID(id)
}
