package order

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/horzu/golang/cart-api/internal/api"
)

func orderToResponse(o Order) *api.Order{
	createdDate := strfmt.Date(o.CreatedAt)

	return &api.Order{
		ID: o.Id,	
		CreatedAt: createdDate,
	}
}

func ordersToResponse(os *[]Order) []*api.Order{
	orders := make([]*api.Order,0)
	for _, order := range *os{
		orders = append(orders, orderToResponse(order))
	}
	return orders
}

func responseToOrder(o *api.Order) *Order{
	return &Order{
	Id: o.ID,
	CreatedAt: time.Time(o.CreatedAt),
	}
}