package productCategory

import (
	"github.com/google/uuid"
	"github.com/horzu/golang/cart-api/internal/api"
	"github.com/horzu/golang/cart-api/internal/models"
)

func categoryToResponse(p *models.ProductCategory) api.Category {
	return api.Category{
		ID:          p.Id,
		Name:       &p.Name,
	}
}

func responseToCategory(a *api.Category) *models.ProductCategory {
	id := uuid.New().String()
	return &models.ProductCategory{
		Id: id,
		Name: *a.Name,
	}
}

