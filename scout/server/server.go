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

	router.GET("/index/:id", controllers.GetIndexController)
	router.POST("/index", controllers.PostIndexController)
	
	router.POST("/ingest", controllers.PostIngestController)
	
	// router.GET("/index", controllers.GetIndexController)
	// router.PUT("/index/:id", controllers.PostIndexController)
	// router.DELETE("/index/:id", controllers.PostIndexController)
	// router.POST("/search", controllers.PostSearchController)
	// router.POST("/backup", controllers.PostIngestController)
	return router
}