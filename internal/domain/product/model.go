package product

import (
	"errors"
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
	CategoryId  string
	Name        string
	Description string
	SKU         string
	Price       float64
	Stock       int64
	IsActive    bool

	Category *category.Category
	// ImageFile []*mediaFile.MediaFile
	// Stock     *stock.Stock
}

func (u *Product) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()
	u.SKU = uuid.New().String()
	u.IsActive = true

	return nil
}

func NewProduct(name string, desc string, stockCount int64, price float64, categoryId string) *Product {
	return &Product{
		Name:        name,
		Description: desc,
		Stock:       stockCount,
		Price:       price,
		CategoryId:  categoryId,
	}
}

func (p *Product) UpdateProduct(name, sku, description, categoryId string, stockQuantity int64, price float64) {
	p.Name = name
	p.SKU = sku
	p.Stock = stockQuantity
	p.Price = price
	p.Description = description
	p.CategoryId = categoryId
}
func (p *Product) UpdateQuantity(changeAmount int64) error {
	err := p.CheckStockExist(changeAmount)
	if err != nil {
		return err
	}
	p.Stock += changeAmount
	return nil
}

func (p *Product) CheckStockExist(changeAmount int64) error {
	if p.Stock+changeAmount < 0 {
		return errors.New("ErrProductStockIsNotEnough")
	}
	return nil
}
