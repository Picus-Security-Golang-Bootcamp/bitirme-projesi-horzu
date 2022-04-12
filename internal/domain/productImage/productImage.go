package productImage

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

	Image *mediafile.MediaFile
}

func (pi *ProductImage) SetImage(f *mediafile.MediaFile) {
	pi.Image = f
}

func (pi *ProductImage) GetImage() *mediafile.MediaFile {
	return pi.Image
}

func (u *ProductImage) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}