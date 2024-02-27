package models

import (
	"gorm.io/gorm"
)

type Index struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null"`
	Description string
	Slug 	  	string `gorm:"type:varchar(100);not null;unique"`
	Status      string `gorm:"type:varchar(20);not null"`
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


func (index *Index) FindAll() ([]Index, error) {
	var indexes []Index
	if err := DB.Find(&indexes).Error; err != nil {
		return []Index{}, err
	}

	return indexes, nil
}

func (index *Index) FindById() (*Index, error) {
	if err := DB.Where("id = ?", index.ID).First(&index).Error; err != nil {
		return &Index{}, err
	}

	return index, nil
}