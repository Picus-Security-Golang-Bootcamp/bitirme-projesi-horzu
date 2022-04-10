package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductImage struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	ProductId int
	ImageId   int

	File *MediaFile
}

func (pi *ProductImage) SetImage(f *MediaFile) {
	pi.File = f
}

func (pi *ProductImage) GetImage() *MediaFile {
	return pi.File
}

func (u *ProductImage) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}