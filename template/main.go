package main 

import (
	"template/server"
)

func main() {
	server := server.Init()
	server.Run(":8000")
}