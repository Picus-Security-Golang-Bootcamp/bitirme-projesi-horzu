package product

import (
	"github.com/google/uuid"
	"github.com/horzu/golang/cart-api/internal/api"
)

func productToResponse(p *Product) api.Product {
	return api.Product{
		ID:          p.Id,
		Price:       &p.Price,
		ProductName: &p.Name,
		Rating:      &p.Rating,
	}
}

func responseToProduct(a *api.Product) *Product {
	id := uuid.New().String()
	return &Product{
		Id:     id,
		Name:   *a.ProductName,
		Price:  *a.Price,
		Rating: *a.Rating,
	}
}
