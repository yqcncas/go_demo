package main

import (
	"bilibili_demo/router"
)

func main() {
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r := router.InitRouter()
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
