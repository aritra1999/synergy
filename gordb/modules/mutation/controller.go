package mutation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MutationBody struct {
	Table string `json:"table"`
}

func MutationController(c *gin.Context) {
	var requestBody MutationBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	table := requestBody.Table
	response, err := MutationProcessor(table)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
