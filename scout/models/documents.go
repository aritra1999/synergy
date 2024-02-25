package models

import (
	"encoding/json"
	"time"
)

type Document struct {
	ID        	string 				`gorm:"primaryKey"`
	Source	  	json.RawMessage 	`gorm:"type:json"`
	Index     	string 
	Type      	string 
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

func (document *Document) Insert() (*Document, error) {

	if err := DB.Create(&document).Error; err != nil {
		return &Document{}, err
	}

	return document, nil
}
