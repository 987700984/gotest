package middleware

import (
	"example.com/v2/controller"
	"github.com/gin-gonic/gin"
	"strconv"
)

// AuthHandler 拦截器
func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cid := c.DefaultPostForm("id", "")
		id, _ := strconv.Atoi(cid)
		if id == 10000 {
			controller.Common{}.ReturnError(c, 2, "苏打粉是")
			//终止
			c.Abort()
		}

		c.Next()
	}
}
