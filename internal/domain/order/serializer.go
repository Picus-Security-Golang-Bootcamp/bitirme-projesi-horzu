package order

import (
	"github.com/horzu/golang/cart-api/internal/api"
)

func orderToResponse(o Order) *api.OrderCompleteOrderResponse {
	// createdDate := strfmt.Date(o.CreatedAt)

	return &api.OrderCompleteOrderResponse{
		OrderID:    o.Id,
		CreatedAt:  o.CreatedAt.String(),
		TotalPrice: o.TotalPrice,
	}
}

func ordersToResponse(os *[]Order) []*api.OrderCompleteOrderResponse {
	orders := make([]*api.OrderCompleteOrderResponse, 0)
	for _, order := range *os {
		orders = append(orders, orderToResponse(order))
	}
	return orders
}

func responseToOrder(o *api.OrderCompleteOrderResponse) *Order {

	return &Order{
		Id:        o.OrderID,
		IsCanceled: false,
	}
}
