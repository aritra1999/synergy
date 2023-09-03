package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func helloController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello world"})
}

func main() {
    r := gin.Default()
    r.GET("/hello", helloController)
    r.Run(":8000")
}