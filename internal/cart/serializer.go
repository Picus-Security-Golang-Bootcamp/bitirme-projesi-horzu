package cart

import (
	"github.com/horzu/golang/cart-api/internal/api"
	"github.com/horzu/golang/cart-api/internal/models"
)

func modelProductToApiProduct(o *models.CartProduct) *api.CartProduct{
	return &api.CartProduct{
		ID: o.Id,	
	}
}

func modelProductsToApiProducts(os *[]models.CartProduct) []*api.CartProduct{
	products := make([]*api.CartProduct,0)
	for _, product := range *os{
		products = append(products, modelProductToApiProduct(&product))
	}
	return products
}
func apiProductToModelProduct(o *api.CartProduct) models.CartProduct{
	return models.CartProduct{
		Id: o.ID,	

	}
}

func apiProductsToModelProducts(os []*api.CartProduct) *[]models.CartProduct{
	products := make([]models.CartProduct,0)
	for _, product := range os{
		products = append(products, apiProductToModelProduct(product))
	}
	return &products
}

func cartToResponse(o *models.Cart) *api.Cart{
	return &api.Cart{
		ID: o.Id,	
		Products: modelProductsToApiProducts(o.Products),		
		TotalPrice: &o.TotalPrice,
	}
}

func cartsToResponse(os *[]models.Cart) []*api.Cart{
	carts := make([]*api.Cart,0)
	for _, cart := range *os{
		carts = append(carts, cartToResponse(&cart))
	}
	return carts
}

func responseToCart(o *api.Cart) *models.Cart{
	return &models.Cart{
	Id: o.ID,
	Products: apiProductsToModelProducts(o.Products),
	TotalPrice: *o.TotalPrice,
	}
}
func responseToCarts(os *[]api.Cart) []*models.Cart{
	carts := make([]*models.Cart,0)
	for _, cart := range *os{
		carts = append(carts, responseToCart(&cart))
	}
	return carts
}

