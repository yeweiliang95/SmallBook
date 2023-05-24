package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()
	engine.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code": 200,
			"msg":  "get test",
		})
	})

	engine.Run()
}
