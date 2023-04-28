package common

import (
	"github.com/gin-gonic/gin"

	"github.com/anaxaim/tui/backend/pkg/model"
	"github.com/anaxaim/tui/backend/pkg/utils"
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

func SetUser(c *gin.Context, user *model.User) {
	if c == nil || user == nil {
		return
	}

	c.Set(UserContextKey, user)
}

func SetRequestInfo(c *gin.Context, ri *utils.RequestInfo) {
	if c == nil || ri == nil {
		return
	}

	c.Set(RequestInfoContextKey, ri)
}
