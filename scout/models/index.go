package models

import (
	"gorm.io/gorm"
)

type Index struct {
	gorm.Model
	Name        string `json:"name";gorm:"type:varchar(100);not null"`
	Description string `json:"description";gorm:"type:varchar(255);not null"`
	Slug 	  	string `json:"slug";gorm:"type:varchar(100);not null;unique"`
	Status      string `json:"status";gorm:"type:varchar(20);not null"`
}


func (index *Index) Insert() (*Index, error) {	
	if err := DB.Create(&index).Error; err != nil {
		return &Index{}, err
	}

	return index, nil
}

func (index *Index) Update() (*Index, error) {
	if err := DB.Save(&index).Error; err != nil {
		return &Index{}, err	
	}

	return index, nil
}

func (index *Index) Delete() (*Index, error) {
	if err := DB.Delete(&index).Error; err != nil {
		return &Index{}, err
	}

	return index, nil
}


func FindAll() ([]Index, error) {
	var indexes []Index
	if err := DB.Find(&indexes).Error; err != nil {
		return []Index{}, err
	}

	return indexes, nil
}

func FindById(id int) (*Index, error) {
	var index Index
	if err := DB.Where("id = ?", id).First(&index).Error; err != nil {
		return &Index{}, err
	}
	
	return &index, nil
}
