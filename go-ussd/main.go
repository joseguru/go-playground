package main

import (
	"example.com/greetings/go-ussd/controllers"
	"example.com/greetings/go-ussd/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook) // new
	r.POST("/ussd", controllers.GetInputs)
	r.Run(":8007")

}
