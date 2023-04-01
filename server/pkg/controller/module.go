package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/anaxaim/tui/server/pkg/common"
	"github.com/anaxaim/tui/server/pkg/model"
	"github.com/anaxaim/tui/server/pkg/service"
)

type ModuleController struct {
	moduleService service.ModuleService
}

func NewModuleController(moduleService service.ModuleService) Controller {
	return &ModuleController{
		moduleService: moduleService,
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

func (m *ModuleController) RegisterRoute(api *gin.RouterGroup) {
	api.GET("/modules", m.List)
	api.POST("/modules", m.Create)
	api.GET("/modules/:id", m.Get)
	api.PUT("/modules/:id", m.Update)
	api.DELETE("/modules/:id", m.Delete)
}

func (m *ModuleController) Name() string {
	return "Module"
}
