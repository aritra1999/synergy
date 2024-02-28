package controllers

import (
	"fmt"
	"net/http"
	"scout/core"
	"scout/models"
	"strconv"

	"github.com/gin-gonic/gin"
)



func GetIndexByIdController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id!"})
	}

	index, err := models.FindById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get index!", "error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"index": index})
}

func GetIndexesController(c *gin.Context) {
	indexes, err := models.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get indexes!", "error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"indexes": indexes})
}

type PostIndexControllerBody struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Slug string `json:"slug,omitempty"`
}

func PostIndexController(c *gin.Context) {
	var err error
	var body PostIndexControllerBody

	if err = c.BindJSON(&body); err != nil {
		fmt.Println("Invalid data!")
		return 
	}

	index := models.Index{
		Name: body.Name, 
		Description: body.Description, 
		Status: "active",
	}

	if body.Slug != "" {
		index.Slug = body.Slug
	} else {
		index.Slug = core.CreateSlug(body.Name)
	}

	var createdIndex *models.Index
	createdIndex, err = index.Insert()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add new index!", "error": err.Error()})
		return
	}


	c.JSON(http.StatusOK, gin.H{"message": "Added new index!", "index": createdIndex })
}