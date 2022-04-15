package cart

import (
	"context"
	"errors"

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
	Create(ctx context.Context, customerId string) (*Cart, error)
	AddItem(ctx context.Context, sku string, cartId string, orderQuantity int64) (string, error)
	UpdateItem(ctx context.Context, id string, cartId string, updateQuantity uint) error
	GetCartItems(ctx context.Context, cartId string) ([]*cartItem.CartItem, error)
	DeleteItem(ctx context.Context, basketId, itemId string) error
	FetchCartByUserId(ctx context.Context,UserID string) (Cart, error)
	ClearBasket(ctx context.Context,cart *Cart) 
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

	basket, _ := service.cartItemRepo.GetItems(ctx, id)
	return basket, nil
}

// Create creates a new cart
func (service *CartService) Create(ctx context.Context, customerId string) (*Cart, error) {
	cart := &Cart{
		UserID: customerId,
		Items:  nil,
	}
	newCart, err := service.cartRepo.Create(ctx, cart)

	if err != nil {
		return nil, errors.New("Service:Failed to create cart")
	}
	return newCart, nil
}

// AddItem adds the product with given amount to given user's cart
func (service *CartService) AddItem(ctx context.Context, sku string, cartId string, orderQuantity int64) (string, error) {
	addedProduct, err := service.productRepo.GetBySku(ctx, sku)
	if err != nil {
		return "", err
	}
	cart, err := service.cartRepo.FindOrCreateByUserID(ctx, cartId)
	if err != nil {
		return "", err
	}
	_, err = service.cartItemRepo.FindByID(ctx, addedProduct.Id, cart.Id)
	if err == nil {
		return "", ErrItemAlreadyInCart
	}
	if addedProduct.Quantity < orderQuantity {
		return "", product.ErrProductInsufficientStock
	}
	if orderQuantity < 0 {
		return "", ErrInvalidOrder
	}
	err = service.cartItemRepo.Create(ctx, cartItem.NewCartItem(addedProduct.Id, cart.Id, uint(orderQuantity)))

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

func (service *CartService) DeleteItem(ctx context.Context, basketId, itemId string) error {

	basket := service.cartRepo.Get(ctx, basketId)
	if basket == nil {
		return errors.New("Service: Get basket error. Basket Id:%s")
	}
	if basket == nil {
		return errors.New("Service: Basket not found")
	}
	if err := service.cartItemRepo.DeleteById(ctx, itemId); err != nil {
		return errors.New("Service: Failed to update basket in data storage.")
	}
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

//FetchCartByUserId it returns cart model for complete
func (service *CartService) FetchCartByUserId(ctx context.Context ,UserID string) (Cart, error) {
	cart, err := service.cartRepo.FindByUserId(ctx, UserID)

	if err != nil {
		return Cart{}, err
	}

	return cart, nil

}

func (service *CartService) ClearBasket(ctx context.Context,cart *Cart) {
	for _, item := range *cart.Items {
		service.cartItemRepo.DeleteById(ctx,item.Id)
	}
}