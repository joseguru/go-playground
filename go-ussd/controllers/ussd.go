package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInputs(c *gin.Context) {
	//sessionId := c.PostForm("sessionId")
	text := c.PostForm("text")

	if IsUserStarting(text) {
		response := GetHomeMenu()
		//fmt.Println(response)
		c.JSON(http.StatusOK, gin.H{"data": response})
	} else {
		response := StateSwitch()
		//fmt.Println(response)

		c.JSON(http.StatusOK, gin.H{"data": response})
	}

}
func IsUserStarting(text string) bool {
	if text == "" {
		return true
	} else {
		return false
	}

}
func GetHomeMenu() string {
	return "home screen"
}

func StateSwitch() string {
	return "state switch"

}
