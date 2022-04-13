package product

import (
	"time"

	"github.com/google/uuid"
	"github.com/horzu/golang/cart-api/internal/domain/category"
	"gorm.io/gorm"
)

type Product struct {
	Id          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	CategoryId  uint
	Name        string
	Slug        string
	Description string
	Price       float64
	Quantity    uint
	Rating      int64
	Weight      float64
	IsActive    bool

	Category  *category.Category
	// ImageFile []*mediaFile.MediaFile
	// Stock     *stock.Stock
}

func (u *Product) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}

func NewProduct(name string, desc string, stockCount uint, price float64, categoryId uint) *Product {
	return &Product{
		Name:       name,
		IsActive:   true,
		Quantity:   stockCount,
		Price:      price,
		CategoryId: categoryId,
	}
}
