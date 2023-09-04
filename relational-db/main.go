package main 

import (
	"relational-db/server"
)

func main() {
	server := server.Init()
	server.Run(":8000")
}