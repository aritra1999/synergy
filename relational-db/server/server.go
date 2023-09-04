package server

import (
	"github.com/gin-gonic/gin"
	// "relational-db/middlewares"
	"relational-db/controllers"
)

func Init() *gin.Engine {
	router := gin.Default()

	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// router.Use(middlewares.AuthMiddleware())

	router.GET("/health", controllers.HealthController)
	router.POST("/query", controllers.PostQueryController)

	return router
}
