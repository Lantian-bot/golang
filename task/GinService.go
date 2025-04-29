package main

import (
	"basic/task4/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Gin Server Start...")
	//创建gin Server
	ginServer := gin.Default()
	// 用户
	controllers.User(ginServer)
	// 文章
	controllers.Post(ginServer)
	// 评论
	controllers.Comment(ginServer)

	// 启动服务
	ginServer.Run(":8080")
}
