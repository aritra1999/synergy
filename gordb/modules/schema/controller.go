package schema

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSchemaController(c *gin.Context) {

	c.String(http.StatusOK, "Working!")
}

func PostSchemaController(c *gin.Context) {

	c.String(http.StatusOK, "Working!")
}

func PutSchemaController(c *gin.Context) {

	c.String(http.StatusOK, "Working!")
}

func DeleteSchemaController(c *gin.Context) {

	c.String(http.StatusOK, "Working!")
}
