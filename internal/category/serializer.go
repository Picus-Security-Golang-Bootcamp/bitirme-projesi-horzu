package category

import (
	"github.com/horzu/golang/cart-api/internal/api"
	"github.com/horzu/golang/cart-api/internal/models"
	"gorm.io/gorm"
)

func categoryToResponse(c models.Category) *api.Category{
	return &api.Category{
		Name: &c.Category_Name,
		Image: &c.Image,
	}
}

func categoriesToResponse(cs *[]models.Category) []*api.Category {
	categories := make([]*api.Category,0)
	
	for _, category := range *cs{
		categories = append(categories, categoryToResponse(category))
	}
	return categories
}

func responseToCategory(a *api.Category) *models.Category {
	return &models.Category{
		Model: gorm.Model{ID: uint(a.ID)},
		Category_Name: *a.Name,
		Image:         *a.Image,
	}
}
