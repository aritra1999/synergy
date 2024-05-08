package query

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Table []map[string]string `json:"insert"`
type QueryBody struct {
	Table string `json:"table"`
}

func QueryController(c *gin.Context) {
	var requestBody QueryBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	table := requestBody.Table
	response, err := QueryProcessor(table)


	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
