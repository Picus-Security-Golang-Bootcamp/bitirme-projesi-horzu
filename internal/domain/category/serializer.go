package category

import (
	"github.com/horzu/golang/cart-api/internal/api"
)

func categoryToResponse(c Category) *api.Category{
	return &api.Category{
		Name: &c.Name,
		Slug: c.Slug,
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
		Slug:         *&a.Slug,
	}
}
