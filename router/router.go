package router

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.New()

	// 创建路由和路由分组，并添加路由处理程序
	router.GET("/api/test", controller.TestFunc)

	// router.GET("/api/get/users", controller.GetUser)

	router.GET("/api/orm/get/user", controller.OrmGetUser)
	router.GET("/api/orm/get/users", controller.OrmGetUsers)

	router.GET("/asc/test", controller.AscTest)

	return router
}
