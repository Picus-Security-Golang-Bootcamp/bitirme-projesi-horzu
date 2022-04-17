package order

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/horzu/golang/cart-api/internal/domain/cart"
	"github.com/horzu/golang/cart-api/internal/domain/order/orderItem"
	"github.com/horzu/golang/cart-api/internal/domain/product"
)

type orderService struct {
	repo           Repository
	orderItemRepo  *orderItem.OrderItemRepository
	cartRepository *cart.CartRepository
	cartService    cart.Service
	productService product.Service
}

type Service interface {
	CompleteOrderWithUserId(ctx context.Context, userId string) error
	GetAll(ctx context.Context, userId string) ([]*Order,error)
	CancelOrder(ctx context.Context, userId, orderId string) error
}

func NewOrderService(orderRepo Repository, orderItemRepo *orderItem.OrderItemRepository, cartService cart.Service, productService product.Service, cartRepository *cart.CartRepository) Service {
	return &orderService{
		repo:           orderRepo,
		orderItemRepo:  orderItemRepo,
		cartService:    cartService,
		productService: productService,
		cartRepository: cartRepository,
	}
}

//CompleteOrderWithUserId crates order from items that is in cart and clear the cart
func (service *orderService) CompleteOrderWithUserId(ctx context.Context, userId string) error {
	cartItems, err := service.cartService.GetAllCartItems(ctx, userId)
	if err != nil {
		return err
	}
	for _, value := range cartItems{

		fmt.Println(value.Product.Price)
	}

	if len(cartItems) == 0 {
		return errors.New("No items in the cart")
	}

	orderedItems := cartItemsToOrderItems(cartItems)

	// Complete Order
	err = service.repo.Create(ctx, NewOrder(userId, orderedItems))

	// Clear cart
	service.cartService.ClearCart(ctx, cartItems)

	return nil
}

//GetAll list all orders of user
func (service *orderService) GetAll(ctx context.Context, userId string) ([]*Order,error) {
	orders, err := service.repo.GetAllByUser(ctx, userId)
	
	if err!=nil{
		return nil, errors.New("No order found")
	}

	return orders, nil
}

//CancelOrder deletes order of the user if 14 days has not passed
func (service *orderService) CancelOrder(ctx context.Context, userId, orderId string) error {
	canceledOrder, err := service.repo.GetByID(ctx, orderId)
	if err != nil {
		return err
	}
	if canceledOrder.UserId != userId {
		return errors.New("No order found")
	}
	if canceledOrder.CreatedAt.Sub(time.Now()).Hours() > days14ToHours {
		return errors.New("Cancel Duration passed")
	}
	canceledOrder.IsCanceled = true
	err = service.repo.Update(ctx, *canceledOrder)
	err = service.repo.DeleteById(ctx ,canceledOrder.Id)

	return err
}