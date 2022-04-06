package product

import (
	"github.com/horzu/golang/cart-api/internal/api"
	"github.com/horzu/golang/cart-api/internal/models"
	"gorm.io/gorm"
)

func productToResponse(p *models.Product) api.Product {
	return api.Product{
		ID:          int64(p.ID),
		Image:       &p.Image,
		Price:       &p.Price,
		ProductName: &p.Product_Name,
		Rating:      &p.Rating,
	}
}

func responseToProduct(a *api.Product) *models.Product {
	return &models.Product{
		Product_Name: *a.ProductName,
		Price:        *a.Price,
		Rating:       *a.Rating,
		Image:        *a.Image,
		Model:        gorm.Model{ID: uint(a.ID)},
	}
}

