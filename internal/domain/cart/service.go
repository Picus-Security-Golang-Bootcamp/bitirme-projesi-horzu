package cart

import (
	"context"
	"errors"
	"fmt"

	"github.com/horzu/golang/cart-api/internal/domain/cart/cartItem"
	"github.com/horzu/golang/cart-api/internal/domain/product"
)

type CartService struct {
	cartRepo     Repository
	productRepo  product.Repository
	cartItemRepo cartItem.Repository
}

type Service interface {
	Get(ctx context.Context, id string) ([]*cartItem.CartItem, error)
	Create(ctx context.Context, customerId string) error
	AddItem(ctx context.Context, sku string, cartId string, orderQuantity int64) (string, error)
	UpdateItem(ctx context.Context, id string, cartId string, updateQuantity uint) error
	GetCartItems(ctx context.Context, cartId string) ([]*cartItem.CartItem, error)
	DeleteItem(ctx context.Context, basketId, itemId string) error
	GetCartByUserId(ctx context.Context, UserID string) (*Cart, error)
	ClearBasket(ctx context.Context, cart *Cart)
}

func NewCartService(r Repository, productRepository product.Repository, cartItemRepository cartItem.Repository) *CartService {
	return &CartService{
		cartRepo:     r,
		productRepo:  productRepository,
		cartItemRepo: cartItemRepository,
	}
}

func (service *CartService) Get(ctx context.Context, id string) ([]*cartItem.CartItem, error) {
	if len(id) == 0 {
		return nil, errors.New("Id cannot be nil or empty")
	}

	cart, err := service.GetCartByUserId(ctx, id)

	fmt.Println(cart.Id)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	cartitems, _ := service.cartItemRepo.GetItems(ctx, cart.Id)
	return cartitems, nil
}

// Create creates a new cart
func (service *CartService) Create(ctx context.Context, customerId string) error {
	cart := &Cart{
		UserID: customerId,
		Items:  nil,
	}
	err := service.cartRepo.Create(ctx, cart)

	if err != nil {
		return errors.New("Service:Failed to create cart")
	}
	return nil
}

// AddItem adds the product with given amount to given user's cart
func (service *CartService) AddItem(ctx context.Context, sku string, cartId string, orderQuantity int64) (string, error) {
	addedProduct, err := service.productRepo.GetBySku(ctx, sku)
	if err != nil {
		return "", err
	}

	_, err = service.cartItemRepo.FindByID(ctx, cartId, addedProduct.Id)
	if err == nil {
		return "", ErrItemAlreadyInCart
	}
	if addedProduct.Stock < orderQuantity {
		return "", product.ErrProductInsufficientStock
	}
	if orderQuantity < 0 {
		return "", ErrInvalidOrder
	}
	err = service.cartItemRepo.Create(ctx, cartItem.NewCartItem(addedProduct.Id, cartId, uint(orderQuantity)))

	return addedProduct.Id, err
}

// UpdateItem updates the amount of product inside given user's cart
func (service *CartService) UpdateItem(ctx context.Context, cartId string, itemId string, updateQuantity uint) error {
	updatedItem, err := service.cartItemRepo.FindByID(ctx, cartId, itemId)
	if err != nil {
		return ErrItemNotExistInCart
	}
	updatedItem.Quantity = updateQuantity
	err = service.cartItemRepo.Update(ctx, updatedItem)

	return err
}

func (service *CartService) DeleteItem(ctx context.Context, cartId, itemId string) error {
	deletedItem, err := service.cartItemRepo.FindByID(ctx, cartId, itemId)

	if err != nil {
		return ErrItemNotExistInCart
	}

	err = service.cartItemRepo.DeleteById(ctx, deletedItem.Id)
		
	if err!=nil{
		return errors.New("Service: Failed to update cart in data storage.")

	}

	// cart := service.cartRepo.Get(ctx, cartId)
	// if cart == nil {
	// 	return errors.New("Service: Get cart error. cart Id:%s")
	// }
	// if cart == nil {
	// 	return errors.New("Service: cart not found")
	// }
	// if err := service.cartItemRepo.DeleteById(ctx, itemId); err != nil {
	// 	return errors.New("Service: Failed to update cart in data storage.")
	// }
	return nil
}

// GetCartItems returns the items inside the given user's cart
func (service *CartService) GetCartItems(ctx context.Context, cartId string) ([]*cartItem.CartItem, error) {
	cart, err := service.cartRepo.FindOrCreateByUserID(ctx, cartId)
	if err != nil {
		return nil, err
	}
	items, err := service.cartItemRepo.GetItems(ctx, cart.Id)
	if err == nil {
		return nil, ErrItemNotExistInCart
	}

	return items, nil
}

//GetCartByUserId it returns cart model for complete
func (service *CartService) GetCartByUserId(ctx context.Context, UserID string) (*Cart, error) {
	cart, err := service.cartRepo.FindByUserId(ctx, UserID)

	if err != nil {
		return nil, err
	}

	return cart, nil

}

func (service *CartService) ClearBasket(ctx context.Context, cart *Cart) {
	for _, item := range *cart.Items {
		service.cartItemRepo.DeleteById(ctx, item.Id)
	}
}
