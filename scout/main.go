package main

import (
	"fmt"
	"scout/models"
	"scout/server"
)

func init() {
	fmt.Println("Initializing server")
	models.ConnectDatabase()
	models.MigrateTables()
}

func main() {
	server := server.Init()
	server.Run(":8000")
}