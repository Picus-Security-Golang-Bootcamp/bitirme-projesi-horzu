package order

import (
	"context"
	"errors"

	"github.com/horzu/golang/cart-api/internal/domain/cart"
	"github.com/horzu/golang/cart-api/internal/domain/order/orderItem"
	"github.com/horzu/golang/cart-api/internal/domain/product"
)

type orderService struct {
	repo           Repository
	orderItemRepo  orderItem.OrderItemRepository
	cartService    cart.Service
	productService product.Service
}

type Service interface {
	CompleteOrderWithUserId(ctx context.Context, userId string) error
}

func NewOrderService(orderRepo *OrderRepository, orderItemRepo *orderItem.OrderItemRepository, cartService cart.Service, productService product.Service) Service {
	return &orderService{
		repo:           orderRepo,
		orderItemRepo:  *orderItemRepo,
		cartService:    cartService,
		productService: productService,
	}
}

//CompleteOrderWithUserId crates order from items that is in basket and clear the basket
func (service *orderService) CompleteOrderWithUserId(ctx context.Context, userId string) error {
	//get user cart
	cart, err := service.cartService.GetCartByUserId(ctx, userId)
	if err != nil {
		return err
	}
	//check cart is empty
	if len(*cart.Items) < 1 {
		return errors.New("basket empty")
	}
	//prepare update product input service
	var productList []product.Product = []product.Product{}
	var orderAmountList []int64 = []int64{}
	for _, item := range *cart.Items {
		productList = append(productList, *item.Product)
		orderAmountList = append(orderAmountList, int64(item.Quantity))
	}
	//update products quantity which are in basket
	errUpdQuant := service.productService.UpdateProductQuantityForOrder(ctx, productList, orderAmountList)
	if errUpdQuant != nil {
		return errUpdQuant
	}

	//create order
	order := NewOrder(*cart)
	_, err1 := service.repo.Create(ctx, order)

	if err1 != nil {
		return err1
	}

	//CLEAR BASKET
	service.cartService.ClearBasket(ctx, cart)

	return nil
}
