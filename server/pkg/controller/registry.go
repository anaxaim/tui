package controller

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/anaxaim/tui/server/pkg/common"
	"github.com/anaxaim/tui/server/pkg/container"
	"github.com/anaxaim/tui/server/pkg/model"
	"github.com/anaxaim/tui/server/pkg/service"
)

type RegistryController struct {
	registryService  service.RegistryService
	terraformService *container.TerraformService
}

func NewRegistryController(registryService service.RegistryService, terraformService *container.TerraformService) Controller {
	return &RegistryController{
		registryService:  registryService,
		terraformService: terraformService,
	}
}

func (r *RegistryController) Import(c *gin.Context) {
	registryContent, err := r.registryService.ImportModuleContentByID(c.Param("id"))
	if err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	common.ResponseSuccess(c, registryContent)
}

func (r *RegistryController) Get(c *gin.Context) {
	registry, err := r.registryService.GetModuleContentByID(c.Param("id"))
	if err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	common.ResponseSuccess(c, registry)
}

func (r *RegistryController) Execute(c *gin.Context) {
	executeCommand := new(model.ExecuteCommand)
	if err := c.BindJSON(executeCommand); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	id := c.Param("id")
	terraformVersion := "latest"

	content, err := r.registryService.GetModuleContentByID(id)
	if err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	workingDir, err := r.terraformService.PrepareWorkingDirectory(content.Content)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	defer os.RemoveAll(workingDir)

	if err := r.terraformService.RunContainer(context.Background(), terraformVersion, workingDir); err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}
	defer r.terraformService.StopContainer(context.Background(), terraformVersion) //nolint: errcheck

	output, err := r.registryService.Execute(terraformVersion, executeCommand.Command)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	common.ResponseSuccess(c, gin.H{"output": string(output)})
}

func (r *RegistryController) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/registry/import/:id", r.Import)
	api.GET("/registry/:id", r.Get)
	api.POST("/registry/execute/:id", r.Execute)
}

func (r *RegistryController) Name() string {
	return "Registry"
}
