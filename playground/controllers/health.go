package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthController(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}