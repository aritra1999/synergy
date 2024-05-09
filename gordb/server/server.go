package server

import (
	"gordb/modules/health"
	"gordb/modules/mutation"
	"gordb/modules/query"
	"gordb/modules/schema"

	"github.com/gin-gonic/gin"
	// "gordb/middlewares"
)

func Init() *gin.Engine {
	router := gin.Default()

	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// router.Use(middlewares.AuthMiddleware())

	router.GET("/health", health.HealthController)

	router.GET("/schema", schema.GetSchemaController)
	router.POST("/schema", schema.PostSchemaController)
	router.PUT("/schema", schema.PutSchemaController)
	router.DELETE("/schema", schema.DeleteSchemaController)

	router.POST("/query", query.QueryController)

	router.POST("/mutation", mutation.MutationController)

	return router
}