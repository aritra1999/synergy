package server

import (
	"scout/controllers"

	"github.com/gin-gonic/gin"
)


func Init() *gin.Engine {
	router := gin.Default()
	
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// router.Use(middlewares.AuthMiddleware())
	
	router.GET("/health", controllers.HealthController)
	router.POST("/ingest", controllers.PostIngestController)
	
	return router
}