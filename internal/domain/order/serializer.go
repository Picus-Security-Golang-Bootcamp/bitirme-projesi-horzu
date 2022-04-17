package order

import (
	"github.com/horzu/golang/cart-api/internal/domain/cart/cartItem"
	"github.com/horzu/golang/cart-api/internal/domain/order/orderItem"
)

func cartItemsToOrderItems(cartItems []cartItem.CartItem) []orderItem.OrderItem {
	orderItems := make([]orderItem.OrderItem, 0)

	for _, cartItem := range cartItems {
		orderItems = append(orderItems, cartItemToOrderItem(cartItem))
	}
	return orderItems
}

func cartItemToOrderItem(cartItem cartItem.CartItem) orderItem.OrderItem  {
	return orderItem.OrderItem {
		Quantity: cartItem.Quantity,
		ProductId: cartItem.ProductId,
		IsCanceled: false,
	}
}