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
	DeletedAt gorm.DeletedAt `gorm:"index"`
	ProductId int
	ImageId   int

	Image *MediaFile
}

func (pi *ProductImage) SetImage(f *MediaFile) {
	pi.Image = f
}

func (pi *ProductImage) GetImage() *MediaFile {
	return pi.Image
}

func (u *ProductImage) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}