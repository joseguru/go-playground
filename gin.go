package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("tes",func(ctx *gin.Context) {
		
	})
	r.Run(":8005") // listen and serve on 0.0.0.0:8080
}
