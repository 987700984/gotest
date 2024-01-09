package controller

import (
	"example.com/v2/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct {
}

func (u UserController) GetUserInfo(c *gin.Context) {
	id := c.Param("id")

	Common{}.ReturnSuccess(c, 0, "success", "user info"+string(id), 1)
}

type Search struct {
	Content string `json:"content"`
	Cid     int    `json:"cid"`
}

func (u UserController) GetList(c *gin.Context) {

	cid := c.PostForm("cid")
	id, _ := strconv.Atoi(cid)
	fmt.Println(id)
	user, _ := model.GetList()
	Common{}.ReturnSuccess(c, 0, id, user, 1)
}

func (u UserController) GetInfo(c *gin.Context) {

	cid := c.PostForm("cid")
	id, _ := strconv.Atoi(cid)
	fmt.Println(id)
	user, _ := model.GetUserTest(id)
	Common{}.ReturnSuccess(c, 0, id, user, 1)
	//content := c.PostForm("content")

	//param := make(map[string]interface{})
	//err := c.BindJSON(&param)
	//if err == nil {
	//	ReturnSuccess(c, 0, param["content"], param["cid"], 1)
	//	return
	//}

	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("捕获异常：", err)
	//	}
	//}()
	//n1 := 1
	//n2 := 0
	//num3 := n1 / n2
	//ReturnError(c, 4004, num3)
	//param := &Search{}
	//err := c.BindJSON(&param)
	//if err == nil {
	//	ReturnSuccess(c, 0, param.Content, param.Cid, 1)
	//	return
	//}
	//ReturnError(c, 4004, "没有相关信息list")
}
func (u UserController) Edit(c *gin.Context) {
	cid := c.PostForm("cid")
	id, _ := strconv.Atoi(cid)
	content := c.PostForm("content")
	model.EditUserTest(id, content)
	Common{}.ReturnSuccess(c, 0, "更新成功", id, 1)

}

func (u UserController) Add(c *gin.Context) {
	content := c.PostForm("content")
	id, err := model.CreateUsetTest(content)
	if err != nil {
		Common{}.ReturnError(c, 1, "删除出错")
	}
	Common{}.ReturnSuccess(c, 0, "新增成功", id, 1)

}

func (u UserController) Del(c *gin.Context) {
	cid := c.PostForm("cid")
	id, _ := strconv.Atoi(cid)
	err := model.DelUserTest(id)
	if err != nil {
		Common{}.ReturnError(c, 1, "删除出错")
	}

	Common{}.ReturnSuccess(c, 0, "删除成功", id, 1)

}
