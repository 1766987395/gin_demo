package main

import (
	"time"

	"net/http"

	"gin/orm"

	"gin/router"

	"gin/config"

	// "gin/db"

	"github.com/gin-gonic/gin"
)

func main() {

	// debug
	gin.SetMode(config.AppMode)

	// 初始化数据库连接
	// db.InitDB()
	// 关闭数据库连接
	// defer db.CloseDB()

	// orm初始化数据库连接
	orm.GormInitDB()
	// orm关闭数据库连接
	defer orm.GormCloseDB()

	// 路由
	r := router.SetupRouter()

	s := &http.Server{
		Addr:           ":" + config.HttpRort,
		Handler:        r,
		ReadTimeout:    time.Duration(config.ReadTimeOut) * time.Second,
		WriteTimeout:   time.Duration(config.WriteTimeOut) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
