package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个默认的路由引擎
	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		// 读取查询参数
		param1 := c.Query("name")
		param2 := c.Query("age")

		// 返回包含查询参数的 JSON 响应
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"param1":  param1,
			"param2":  param2,
		})
	})

	// 路由组
	userGroup := engine.Group("/user")
	{
		userGroup.GET("/info", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "user info",
			})
		})
	}

	engine.Run(":8080")
}
