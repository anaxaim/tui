package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/anaxaim/tui/pkg/common"
	"github.com/anaxaim/tui/pkg/utils"
)

func RequestInfoMiddleware(resolver utils.RequestInfoResolver) gin.HandlerFunc {
	return func(c *gin.Context) {
		ri, err := resolver.NewRequestInfo(c.Request)
		if err != nil {
			common.ResponseFailed(c, http.StatusBadRequest, err)
			c.Abort()
			return
		}

		common.SetRequestInfo(c, ri)

		c.Next()
	}
}
