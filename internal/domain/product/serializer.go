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

func requestToProduct(p *api.ProductCreateProductRequest) *Product{
	return &Product{
		CategoryId: *p.CategoryID,
		Name: *p.Name,
		Description: *p.Description,
		Price: *p.Price,
		Stock: int64(*p.Stock),
	}
}

func responseToProduct(a *api.ProductCreateProductRequest) *Product {
	return &Product{
		Name:        *a.Name,
		Description: *a.Description,
		Price:       *a.Price,
		Stock:       int64(*a.Stock),
		CategoryId:  *a.CategoryID,
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

func requestToUpdateProduct(p *api.ProductUpdateProductRequest) *Product{
	return &Product{
		CategoryId: p.CategoryID,
		Name: p.Name,
		Description: p.Description,
		Price: p.Price,
		Stock: int64(p.Stock),
		SKU: p.Sku,
	}
}