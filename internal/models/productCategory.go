package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductCategory struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Name      string
	Slug      string
	IsActive  bool
}

func (u *ProductCategory) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}