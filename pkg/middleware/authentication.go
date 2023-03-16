package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/anaxaim/tui/pkg/authentication"
	"github.com/anaxaim/tui/pkg/common"
	"github.com/anaxaim/tui/pkg/repository"
)

func AuthenticationMiddleware(jwtService *authentication.JWTService, userRepo repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, _ := getTokenFromAuthorizationHeader(c)
		if token == "" {
			token, _ = getTokenFromCookie(c)
		}

		user, _ := jwtService.ParseToken(token)
		if user != nil {
			user, err := userRepo.GetUserByName(user.Name)
			if err != nil {
				common.ResponseFailed(c, http.StatusInternalServerError, fmt.Errorf("failed to get user"))
				c.Abort()
				return
			}
			common.SetUser(c, user)
		}

		c.Next()
	}
}

func getTokenFromCookie(c *gin.Context) (string, error) {
	return c.Cookie("token")
}

func getTokenFromAuthorizationHeader(c *gin.Context) (string, error) {
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		return "", nil
	}

	token := strings.Fields(auth)
	if len(token) != 2 || strings.ToLower(token[0]) != "bearer" || token[1] == "" {
		return "", fmt.Errorf("authorization header invaild")
	}

	return token[1], nil
}
