package controllers

import (
	"net/http"
	"relationalDb/processor"

	"github.com/gin-gonic/gin"
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
	response, err := processor.QueryProcessor(query)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
