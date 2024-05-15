package schema

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTableController(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

func PostTableController(c *gin.Context) {
	var tables []Table

	if err := c.ShouldBindJSON(&tables); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tableResponse, err := AddTables(tables)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tables": tableResponse})
}

func PutTableController(c *gin.Context) {

	c.String(http.StatusOK, "Working!")
}

func DeleteTableController(c *gin.Context) {

	c.String(http.StatusOK, "Working!")
}
