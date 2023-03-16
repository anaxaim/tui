package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/anaxaim/tui/pkg/authentication"
	"github.com/anaxaim/tui/pkg/common"
	"github.com/anaxaim/tui/pkg/model"
	"github.com/anaxaim/tui/pkg/service"
)

type AuthController struct {
	userService service.UserService
	jwtService  *authentication.JWTService
}

func NewAuthController(userService service.UserService, jwtService *authentication.JWTService) Controller {
	return &AuthController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (ac *AuthController) Login(c *gin.Context) {
	auser := new(model.AuthUser)
	if err := c.BindJSON(auser); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	var user *model.User
	var err error

	user, err = ac.userService.Auth(auser)
	if err != nil {
		common.ResponseFailed(c, http.StatusUnauthorized, err)
		return
	}

	token, err := ac.jwtService.CreateToken(user)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
		return
	}
	if auser.SetCookie {
		c.SetCookie(common.CookieTokenName, token, 3600*24, "/", "", true, true)
		c.SetCookie(common.CookieLoginUser, string(userJSON), 3600*24, "/", "", true, false)
	}

	common.ResponseSuccess(c, model.JWTToken{
		Token:    token,
		Describe: "set token in Authorization Header, [Authorization: Bearer {token}]",
	})
}

func (ac *AuthController) Logout(c *gin.Context) {
	c.SetCookie(common.CookieTokenName, "", -1, "/", "", true, true)
	c.SetCookie(common.CookieLoginUser, "", -1, "/", "", true, false)
	common.ResponseSuccess(c, nil)
}

func (ac *AuthController) Register(c *gin.Context) {
	createdUser := new(model.CreatedUser)
	if err := c.BindJSON(createdUser); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	user := createdUser.GetUser()
	if err := ac.userService.Validate(user); err != nil {
		common.ResponseFailed(c, http.StatusBadRequest, err)
		return
	}

	user, err := ac.userService.Create(user)
	if err != nil {
		common.ResponseFailed(c, http.StatusInternalServerError, err)
	}

	common.ResponseSuccess(c, user)
}

func (ac *AuthController) RegisterRoute(api *gin.RouterGroup) {
	api.POST("/auth/token", ac.Login)
	api.DELETE("/auth/token", ac.Logout)
	api.POST("/auth/user", ac.Register)
}

func (ac *AuthController) Name() string {
	return "Authentication"
}
