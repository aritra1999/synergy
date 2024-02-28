package relationalDb

import (
	"relationalDb/server"
)

func main() {
	server := server.Init()
	server.Run(":8000")
}
