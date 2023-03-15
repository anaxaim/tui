package common

import (
	"github.com/gin-gonic/gin"

	"github.com/anaxaim/tui/pkg/model"
)

func GetUser(c *gin.Context) *model.User {
	if c == nil {
		return nil
	}

	val, ok := c.Get(UserContextKey)
	if !ok {
		return nil
	}

	user, ok := val.(*model.User)
	if !ok {
		return nil
	}

	return user
}
