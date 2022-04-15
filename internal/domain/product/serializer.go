package product

import (
	"github.com/horzu/golang/cart-api/internal/api"
)

func productToResponse(p *Product) api.Product {
	return api.Product{
		Sku:        &p.SKU,
		Name:       &p.Name,
		Desc:       &p.Description,
		Price:      &p.Price,
		StockCount: &p.Stock,
		CategoryID: &p.CategoryId,
	}
}

func responseToProduct(a *api.Product) *Product {
	return &Product{
		SKU:         *a.Sku,
		Name:        *a.Name,
		Description: *a.Desc,
		Price:       *a.Price,
		Stock:    *a.StockCount,
		CategoryId:  *a.CategoryID,
	}
}
