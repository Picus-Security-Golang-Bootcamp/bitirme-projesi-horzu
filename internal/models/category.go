package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model   `bson:"id"`
	Category_Name string `json:"category_name"`
	Image        string `json:"image"`
}

func (Category) TableName() string{
	// default table name
	return "categories"
}