package controllers

import (
	"fmt"
	"net/http"
	"scout/core"

	"github.com/gin-gonic/gin"
)

type PostIngestControllerBody struct {
	Data []core.Entry `json:"data"`
}


func PostIngestController(c *gin.Context) {
	var body PostIngestControllerBody
	if err:= c.BindJSON(&body); err != nil {
		fmt.Println("Invalid data!")
		return 
	}

	ingested := 0
	total := len(body.Data)

	for _, entry := range body.Data {
		if err := core.Ingest(entry); err == nil {
			ingested++
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Ingested %d / %d", ingested, total)})
}

 