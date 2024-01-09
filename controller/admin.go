package controller

import (
	"example.com/v2/model"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type AdminController struct{}

func (AdminController) Register(c *gin.Context) {
	//接受用户名，密码
	username := c.PostForm("username")
	pwd := c.PostForm("password")
	if username == "" || pwd == "" {
		Common{}.ReturnError(c, 1, "参数错误")
	}

}

func (AdminController) Login(c *gin.Context) {

	//接受用户名，密码
	username := c.PostForm("username")
	pwd := c.PostForm("password")
	if username == "" || pwd == "" {
		Common{}.ReturnError(c, 1, "参数错误")
		return
	}
	var param model.Admin
	err := c.ShouldBind(&param)
	fmt.Println(param.Password)
	admin, err := model.GetAdminTest(username)
	if err != nil {
		Common{}.ReturnError(c, 1, err)
		return
	}
	session := sessions.Default(c)
	fmt.Println(session)
	session.Set("login:"+admin.Username, admin.Id)
	session.Save()

	//测试redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // 密码
		DB:       0,                // 使用默认数据库
	})
	defer client.Close()
	// 设置key
	err = client.Set(c, "name", "john", 0).Err()
	if err != nil {
		panic(err)
	}

	// 获取key
	val, err := client.Get(c, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name", val)
	Common{}.ReturnSuccess(c, 0, "登录成功", param, 1)

}

func (AdminController) redis(c *gin.Context) {

	//测试redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // 密码
		DB:       0,                // 使用默认数据库
	})
	defer client.Close()
	// 设置key
	err := client.Set(c, "name", "john", 0).Err()
	if err != nil {
		panic(err)
	}

	// 获取key
	val, err := client.Get(c, "name").Result()
	if err != nil {
		panic(err)
	}

	pubsub := client.Subscribe(c, "mychannel")
	defer pubsub.Close()

	// 处理订阅接收到的消息
	for {
		msg, err := pubsub.ReceiveMessage(c)
		if err != nil {
			return
		}

		fmt.Println(msg.Channel, msg.Payload)
	}

	fmt.Println("name", val)
	Common{}.ReturnSuccess(c, 0, "登录成功", val, 1)

}
