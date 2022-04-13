package category

import (
	"github.com/horzu/golang/cart-api/internal/api"
)

func categoryToResponse(c Category) *api.Category{
	return &api.Category{
		Name: &c.Name,
		Description: c.Description,
	}
}

func categoriesToResponse(cs *[]Category) []*api.Category {
	categories := make([]*api.Category,0)
	
	for _, category := range *cs{
		categories = append(categories, categoryToResponse(category))
	}
	return categories
}

func responseToCategory(a *api.Category) *Category {
	return &Category{
		Id: a.ID,
		Name: *a.Name,
		Description:         *&a.Description,
	}
}
