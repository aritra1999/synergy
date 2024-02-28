package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	Source	  	json.RawMessage 	`gorm:"type:json"`
	Index     	string 
	Type      	string 
	CreatedAt 	time.Time
}

func (document *Document) Insert() (*Document, error) {

	if err := DB.Create(&document).Error; err != nil {
		return &Document{}, err
	}

	return document, nil
}
