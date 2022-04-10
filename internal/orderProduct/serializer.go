package orderProduct

import (
	"github.com/google/uuid"
	"github.com/horzu/golang/cart-api/internal/api"
	"github.com/horzu/golang/cart-api/internal/models"
)

func productToResponse(p *models.Product) api.Product {
	return api.Product{
		ID:          p.Id,
		Price:       &p.Price,
		ProductName: &p.Name,
		Rating:      &p.Rating,
	}
}

func responseToProduct(a *api.Product) *models.Product {
	id := uuid.New().String()
	return &models.Product{
		Id:     id,
		Name:   *a.ProductName,
		Price:  *a.Price,
		Rating: *a.Rating,
	}
}
