package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `json:"name" form:"name"`
	Item []Item `gorm:"foreignkey:CategoryID"`
}
