package playground

import (
	"playground/server"
)

func main() {
	server := server.Init()
	server.Run(":8000")
}
