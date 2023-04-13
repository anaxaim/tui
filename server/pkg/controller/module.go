package controller

import (
	"context"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/anaxaim/tui/server/pkg/common"
	"github.com/anaxaim/tui/server/pkg/container"
	"github.com/anaxaim/tui/server/pkg/model"
	"github.com/anaxaim/tui/server/pkg/service"
)

type ModuleController struct {
	moduleService    service.ModuleService
	terraformService *container.TerraformService
}

func NewModuleController(moduleService service.ModuleService, terraformService *container.TerraformService) Controller {
	return &ModuleController{
		moduleService:    moduleService,
		terraformService: terraformService,
	}
}

func (m *ModuleController) List(c *gin.Context) {
	modules, err := m.moduleService.List()
	if err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	common.ResponseSuccess(c, modules)
}

func (m *ModuleController) Get(c *gin.Context) {
	module, err := m.moduleService.Get(c.Param("id"))
	if err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	common.ResponseSuccess(c, module)
}

func (m *ModuleController) Create(c *gin.Context) {
	module := new(model.TerraformModule)
	if err := c.BindJSON(module); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	if err := m.moduleService.Validate(module); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	module, err := m.moduleService.Create(module)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	common.ResponseSuccess(c, module)
}

func (m *ModuleController) Update(c *gin.Context) {
	newModule := new(model.TerraformModule)

	if err := c.BindJSON(newModule); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	logrus.Infof("get update module: %#v", newModule)

	module, err := m.moduleService.Update(c.Param("id"), newModule)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	common.ResponseSuccess(c, module)
}

func (m *ModuleController) Delete(c *gin.Context) {
	if err := m.moduleService.Delete(c.Param("id")); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	common.ResponseSuccess(c, nil)
}

func (m *ModuleController) Import(c *gin.Context) {
	moduleID := c.Param("id")
	terraformVersion := "latest"

	workingTmpDir := filepath.Join(os.TempDir(), moduleID)
	if err := os.MkdirAll(workingTmpDir, 0o777); err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	defer os.RemoveAll(workingTmpDir)

	module, err := m.moduleService.ImportModuleContent(moduleID, workingTmpDir)
	if err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	if err := m.terraformService.RunContainerWithVolume(context.Background(), terraformVersion, workingTmpDir, moduleID); err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}
	// defer m.terraformService.StopContainer(context.Background(), terraformVersion)

	common.ResponseSuccess(c, module)
}

func (m *ModuleController) Execute(c *gin.Context) {
	executeCommand := new(model.ExecuteCommand)
	if err := c.BindJSON(executeCommand); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	id := c.Param("id")
	terraformVersion := "latest"

	if err := m.terraformService.RunContainer(context.Background(), terraformVersion); err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}
	// defer m.terraformService.StopContainer(context.Background(), terraformVersion)

	output, err := m.moduleService.Execute(terraformVersion, executeCommand.Command, id)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	common.ResponseSuccess(c, gin.H{"output": string(output)})
}

func (m *ModuleController) RegisterRoute(api *gin.RouterGroup) {
	api.GET("/modules", m.List)
	api.POST("/modules", m.Create)
	api.GET("/modules/:id", m.Get)
	api.PUT("/modules/:id", m.Update)
	api.DELETE("/modules/:id", m.Delete)
	api.POST("/modules/import/:id", m.Import)
	api.POST("/modules/execute/:id", m.Execute)
}

func (m *ModuleController) Name() string {
	return "Module"
}
