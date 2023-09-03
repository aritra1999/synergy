package main 

import (
	"postgresql/server"
)

func main() {
	server := server.Init()
	server.Run(":8000")
}