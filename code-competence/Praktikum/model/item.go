package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	CategoryID  uint     `json:"category_id" form:"category_id"`
	Name        string   `json:"name" form:"name"`
	Description string   `json:"description" form:"description"`
	Stock       int      `json:"stock" form:"stock"`
	Price       int      `json:"price" form:"price"`
	Category    Category `json:"-" gorm:"foreignKey:CategoryID"`
}
