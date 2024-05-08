package main

import "gordb/server"

func main() {
	server := server.Init()
	server.Run(":8000")
}
