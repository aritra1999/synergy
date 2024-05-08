package query

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueryController(c *gin.Context) {
	var query Query

	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if valid := ValidateQuery(query); !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query"})
		return
	} 
		
	response, err := QueryProcessor(query)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
