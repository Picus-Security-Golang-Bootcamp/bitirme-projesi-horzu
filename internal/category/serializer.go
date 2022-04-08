package category

import (
	"github.com/horzu/golang/cart-api/internal/api"
	"github.com/horzu/golang/cart-api/internal/models"
	"gorm.io/gorm"
)

func categoryToResponse(p *[]models.Category) []api.Category {
	var allItems []api.Category
	
	for _, category := range *p{
		allItems = append(allItems, api.Category{
			Name: &category.Category_Name,
		})
	}
	return allItems
}

func responseToCategory(a *api.Category) *models.Category {
	return &models.Category{
		Model: gorm.Model{ID: uint(a.ID)},
		Category_Name: *a.Name,
		Image:         *a.Image,
	}
}
