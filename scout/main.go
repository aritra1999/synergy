package main

import (
	"scout/server"
)

func main() {
	server := server.Init()
	server.Run(":8000")
}