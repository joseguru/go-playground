package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UssdInput struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}


func index(c *gin.Context) {
	var input UssdInput
	if err:= c.ShouldBindJSON(&input); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
func isUserStarting(){
	
}
