package router

import (
	"example.com/v2/controller"
	"example.com/v2/middleware"
	"example.com/v2/pkg/logger"
	"github.com/gin-contrib/sessions"
	session_redis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	})
	//重写日志
	r.Use(gin.LoggerWithConfig(logger.LoggerToFile()))

	//抛出错误，写入日志
	r.Use(logger.Recover)

	//使用session
	store, _ := session_redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	//全局中间键

	r.Use(middleware.AuthHandler())
	user := r.Group("/user")
	{
		user.GET("/info/:id", controller.UserController{}.GetInfo)
		user.POST("/list", controller.UserController{}.GetList)
		user.POST("/add", controller.UserController{}.Add)
		user.POST("/edit", controller.UserController{}.Edit)
		user.POST("/del", controller.UserController{}.Del)
	}
	admin := r.Group("/admin")
	{
		//局部中间件
		admin.POST("/login", middleware.AuthHandler(), controller.AdminController{}.Login)
	}
	image := r.Group("/image")
	{
		image.POST("/upload", controller.ImageController{}.Upload)
	}
	co := r.Group("/co")
	{
		co.POST("/test", controller.CoController{}.Test)
	}
	order := r.Group("/order")
	{
		order.POST("/list", controller.OrderController{}.GetList)
	}

	return r
}
