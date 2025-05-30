package main

import (
	"log"
	"synergy/rdb"

	"github.com/gin-gonic/gin"
)

func init() {
	log.Println("Initializing rdb")
}

func main() {
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Use(gin.Logger())

	router.GET("/ping", rdb.PingController)

	router.Run(":8000")
	return

}
