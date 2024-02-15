package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostIngestControllerBody struct {
	Data []Entry `json:"data"`
}

type Entry struct {
	Entry map[string]interface{} 
}


func PostIngestController(c *gin.Context) {
	var body PostIngestControllerBody
	if err:= c.BindJSON(&body); err != nil {
		fmt.Println("Invalid data!")
		return 
	}

	c.JSON(http.StatusOK, gin.H{"message": "Works fine"})
}
