package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	Id          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	CategoryId  int
	Name        string
	Slug        string
	Description *string
	Price       float64
	Quantity    int
	Rating      int64
	Weight      float64
	IsActive    bool

	Category  *Category
	ImageFile []*ProductImage
	Stock     *Stock
}

func (u *Product) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}
