package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/anaxaim/tui/server/pkg/common"
	"github.com/anaxaim/tui/server/pkg/model"
	"github.com/anaxaim/tui/server/pkg/service"
)

type CredentialController struct {
	credentialService service.CredentialService
}

func NewCredentialController(credentialService service.CredentialService) Controller {
	return &CredentialController{
		credentialService: credentialService,
	}
}

func (cred *CredentialController) List(c *gin.Context) {
	credentials, err := cred.credentialService.List()
	if err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	common.ResponseSuccess(c, credentials)
}

func (cred *CredentialController) Get(c *gin.Context) {
	credential, err := cred.credentialService.Get(c.Param("id"))
	if err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	common.ResponseSuccess(c, credential)
}

func (cred *CredentialController) Create(c *gin.Context) {
	credential := new(model.Credential)
	if err := c.BindJSON(credential); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	credential, err := cred.credentialService.Create(credential)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	common.ResponseSuccess(c, credential)
}

func (cred *CredentialController) Delete(c *gin.Context) {
	if err := cred.credentialService.Delete(c.Param("id")); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	common.ResponseSuccess(c, nil)
}

func (cred *CredentialController) RegisterRoute(api *gin.RouterGroup) {
	api.GET("/credentials", cred.List)
	api.POST("/credentials", cred.Create)
	api.GET("/credentials/:id", cred.Get)
	api.DELETE("/credentials/:id", cred.Delete)
}

func (cred *CredentialController) Name() string {
	return "Credential"
}
