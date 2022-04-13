package productImage

import (
	"time"

	"github.com/google/uuid"
	"github.com/horzu/golang/cart-api/internal/domain/mediaFile"
	"gorm.io/gorm"
)

type ProductImage struct {
	Id        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	ProductId int
	ImageId   int

	Image *mediaFile.MediaFile
}

func (pi *ProductImage) SetImage(f *mediaFile.MediaFile) {
	pi.Image = f
}

func (pi *ProductImage) GetImage() *mediaFile.MediaFile {
	return pi.Image
}

func (u *ProductImage) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}