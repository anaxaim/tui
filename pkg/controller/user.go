package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/anaxaim/tui/pkg/common"
	"github.com/anaxaim/tui/pkg/model"
	"github.com/anaxaim/tui/pkg/service"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) Controller {
	return &UserController{
		userService: userService,
	}
}

func (u *UserController) List(c *gin.Context) {
	users, err := u.userService.List()
	if err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}
	common.ResponseSuccess(c, users)
}

func (u *UserController) Get(c *gin.Context) {
	user, err := u.userService.Get(c.Param("username"))
	if err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}
	common.ResponseSuccess(c, user)
}

func (u *UserController) Create(c *gin.Context) {
	user := new(model.User)
	if err := c.BindJSON(user); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	if err := u.userService.Validate(user); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	user, err := u.userService.Create(user)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
	}

	common.ResponseSuccess(c, user)
}

func (u *UserController) Update(c *gin.Context) {
	user := common.GetUser(c)
	if user == nil {
		common.ResponseFailed(c, http.StatusForbidden, nil)
		return
	}

	newUser := new(model.User)
	if err := c.BindJSON(newUser); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}
	logrus.Infof("get update user: %#v", newUser)

	user, err := u.userService.Update(c.Param("username"), newUser)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	common.ResponseSuccess(c, user)
}

func (u *UserController) Delete(c *gin.Context) {
	user := common.GetUser(c)
	if user == nil {
		common.ResponseFailed(c, http.StatusForbidden, nil)
		return
	}

	if err := u.userService.Delete(c.Param("username")); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	common.ResponseSuccess(c, nil)
}

func (u *UserController) RegisterRoute(api *gin.RouterGroup) {
	api.GET("/users", u.List)
	api.POST("/users", u.Create)
	api.GET("/users/:username", u.Get)
	api.PUT("/users/:username", u.Update)
	api.DELETE("/users/:username", u.Delete)
}

func (u *UserController) Name() string {
	return "User"
}
