package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/anaxaim/tui/server/pkg/common"
	"github.com/anaxaim/tui/server/pkg/model"
	"github.com/anaxaim/tui/server/pkg/service"
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
	user, err := u.userService.Get(c.Param("name"))
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
		if errors.Is(err, service.ErrUserAlreadyExists) {
			common.ResponseFailed(c, http.StatusConflict, err)
			return
		}
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	common.ResponseSuccess(c, user)
}

func (u *UserController) Update(c *gin.Context) {
	newUser := new(model.User)
	if err := c.BindJSON(newUser); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}
	logrus.Infof("get update user: %#v", newUser)

	user, err := u.userService.Update(c.Param("name"), newUser)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	common.ResponseSuccess(c, user)
}

func (u *UserController) Delete(c *gin.Context) {
	if err := u.userService.Delete(c.Param("name")); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	common.ResponseSuccess(c, nil)
}

func (u *UserController) RegisterRoute(api *gin.RouterGroup) {
	api.GET("/users", u.List)
	api.POST("/users", u.Create)
	api.GET("/users/:name", u.Get)
	api.PUT("/users/:name", u.Update)
	api.DELETE("/users/:name", u.Delete)
}

func (u *UserController) Name() string {
	return "User"
}
