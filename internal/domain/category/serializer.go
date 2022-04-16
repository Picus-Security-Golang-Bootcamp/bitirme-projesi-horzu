package category

import (
	"github.com/horzu/golang/cart-api/internal/api"
)

func categoryToResponse(c Category) *api.CategoryListCategoryResponse {
	return &api.CategoryListCategoryResponse{
		Name:        c.Name,
		Description: c.Description,
	}
}

func categoriesToResponse(cs *[]Category) []*api.CategoryListCategoryResponse {
	categories := make([]*api.CategoryListCategoryResponse, 0)

	for _, category := range *cs {
		categories = append(categories, categoryToResponse(category))
	}
	return categories
}

func responseToCategory(a *api.CategoryListCategoryResponse) *Category {
	return &Category{
		Name:        a.Name,
		Description: *&a.Description,
	}
}
