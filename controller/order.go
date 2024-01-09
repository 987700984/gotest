package controller

import "github.com/gin-gonic/gin"

type OrderController struct {
}

func (o OrderController) GetList(c *gin.Context) {
	Common{}.ReturnError(c, 0, "没有相关order")
}
