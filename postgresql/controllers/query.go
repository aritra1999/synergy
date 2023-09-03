package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Body struct {
	Query string `json:"query" binding:"required"`
}

func PostQueryController(c *gin.Context) {
	var requestBody Body
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := requestBody.Query

	c.IndentedJSON(http.StatusOK, gin.H{"query": query})
}
