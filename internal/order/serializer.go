package order

import (
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/horzu/golang/cart-api/internal/api"
	"github.com/horzu/golang/cart-api/internal/models"
	"gorm.io/gorm"
)

func orderToResponse(o models.Order) *api.Order{
	createdDate := strfmt.Date(o.CreatedAt)

	return &api.Order{
		ID: int64(o.ID),	
		Discount: &o.Discount,
		OrderCart: nil,
		CreatedAt: createdDate,
	}
}

func ordersToResponse(os *[]models.Order) []*api.Order{
	orders := make([]*api.Order,0)
	
	for _, order := range *os{
		orders = append(orders, orderToResponse(order))
	}
	return orders
}

func responseToOrder(o *api.Order) *models.Order{
	return &models.Order{
	Model: gorm.Model{ID: uint(o.ID)},
	Ordered_At: time.Time(o.CreatedAt),
	Discount: *o.Discount,
	}
}