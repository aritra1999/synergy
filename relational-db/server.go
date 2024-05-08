package main

import (
	"relationalDb/health"
	"relationalDb/mutation"
	"relationalDb/query"

	"github.com/gin-gonic/gin"
	// "relationalDb/middlewares"
)

func Init() *gin.Engine {
	router := gin.Default()

	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// router.Use(middlewares.AuthMiddleware())

	router.GET("/health", health.HealthController)
	router.POST("/query", query.QueryController)
	router.POST("/mutation", mutation.MutationController)

	return router
}
