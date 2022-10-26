package Init

import (
	"github.com/gin-gonic/gin"
	"metadata/util"
)

func GinRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		util.ResponseSuccess(c, "pong")
	})
}
