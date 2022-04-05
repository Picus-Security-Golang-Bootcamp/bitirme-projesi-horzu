package order

import (
	"github.com/go-openapi/strfmt"
	"github.com/horzu/golang/cart-api/internal/api"
	"github.com/horzu/golang/cart-api/internal/models"
)

func OrderToResponse(o models.Order) *api.Order{
	createdDate := strfmt.Date(o.CreatedAt)

	return &api.Order{
		ID: int64(o.ID),	
		Discount: &o.Discount,
		OrderCart: nil,
		CreatedAt: createdDate,
	}
}

func OrdersToResponse(os *[]models.Order) []*api.Order{
	orders := make([]*api.Order,0)
	
	for _, order := range *os{
		orders = append(orders, OrderToResponse(order))
	}
	return orders
}