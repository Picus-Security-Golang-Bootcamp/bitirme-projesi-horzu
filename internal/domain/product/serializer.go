package product

import (
	"github.com/horzu/golang/cart-api/internal/api"
)

func productToResponse(p *Product) api.ProductGetProductResponse {
	return api.ProductGetProductResponse{
		Sku:        p.SKU,
		Name:       p.Name,
		Desc:       p.Description,
		Price:      p.Price,
		Stock:      p.Stock,
		CategoryID: p.CategoryId,
	}
}

func responseToProduct(a *api.ProductCreateProductRequest) *Product {
	return &Product{
		Name:        a.Name,
		Description: a.Description,
		Price:       a.Price,
		Stock:       int64(a.Stock),
		CategoryId:  a.CategoryID,
	}
}
func responseToUpdateProduct(a *api.ProductUpdateProductRequest) *Product {
	return &Product{
		Name:        a.Name,
		Description: a.Description,
		Price:       a.Price,
		Stock:       int64(a.Stock),
		CategoryId:  a.CategoryID,
	}
}
