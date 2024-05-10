package main

import (
	"fmt"
	"gordb/server"
)

func init() {
	fmt.Println("Initializing server...")
	server.CheckDirs()
}

func main() {
	server := server.Init()
	server.Run(":8000")
}
