package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/anaxaim/tui/server/pkg/common"
	"github.com/anaxaim/tui/server/pkg/service"
)

type RegistryController struct {
	registryService service.RegistryService
}

func NewRegistryController(registryService service.RegistryService) Controller {
	return &RegistryController{
		registryService: registryService,
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

func (r *RegistryController) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/registry/import/:id", r.Import)
	api.GET("/registry/:id", r.Get)
}

func (r *RegistryController) Name() string {
	return "Registry"
}
