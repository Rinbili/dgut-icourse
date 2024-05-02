package main

import (
	"dgut-icourse/app/routers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	// 初始化路由
	routers.InitRouter(r)
	// 运行
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
