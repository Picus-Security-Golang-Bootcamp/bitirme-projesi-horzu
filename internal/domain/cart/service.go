package cart

import (
	"github.com/horzu/golang/cart-api/internal/domain/cartItem"
	"github.com/horzu/golang/cart-api/internal/domain/product"
)

type CartService struct {
	cartRepo          CartRepository
	productRepository product.ProductRepository
	cartItemRepo      cartItem.CartItemRepository
}

func NewCartService(r CartRepository, productRepository product.ProductRepository, cartItemRepository cartItem.CartItemRepository) *CartService {
	return &CartService{
		cartRepo:          r,
		productRepository: productRepository,
		cartItemRepo:      cartItemRepository,
	}
}

// AddItem adds the product with given amount to given user's cart
func (service *CartService) AddItem(sku string, userID string, orderQuantity uint) error {
	addedProduct, err := service.productRepository.GetBySku(sku)
	if err != nil {
		return err
	}
	cart, err := service.cartRepo.FindOrCreateByUserID(userID)
	if err != nil {
		return err
	}
	_, err = service.cartItemRepo.FindByID(addedProduct.Id, cart.Id)
	if err == nil {
		return ErrItemAlreadyInCart
	}
	if addedProduct.Quantity < orderQuantity {
		return product.ErrProductInsufficientStock
	}
	if orderQuantity < 0 {
		return ErrInvalidOrder
	}
	err = service.cartItemRepo.Create(cartItem.NewCartItem(addedProduct.Id, cart.Id, orderQuantity))

	return err
}

// UpdateItem updates the amount of product inside given user's cart
func (service *CartService) UpdateItem(sku string, userID string, updateQuantity uint) error {
	updatedProduct, err := service.productRepository.GetBySku(sku)
	if err != nil {
		return err
	}
	cart, err := service.cartRepo.FindOrCreateByUserID(userID)
	if err != nil {
		return err
	}
	updatedItem, err := service.cartItemRepo.FindByID(updatedProduct.Id, cart.Id)
	if err == nil {
		return ErrItemNotExistInCart
	}
	if updatedProduct.Quantity+updatedItem.Quantity < updateQuantity {
		return product.ErrProductInsufficientStock
	}
	if updateQuantity < 0 {
		return ErrInvalidOrder
	}
	updatedItem.Quantity = updateQuantity
	err = service.cartItemRepo.Update(updatedItem)

	return err
}

// GetCartItems returns the items inside the given user's cart
func (service *CartService) GetCartItems(userID string) ([]cartItem.CartItem, error) {
	cart, err := service.cartRepo.FindOrCreateByUserID(userID)
	if err != nil {
		return nil, err
	}
	items, err := service.cartItemRepo.GetItems(cart.Id)
	if err == nil {
		return nil, ErrItemNotExistInCart
	}

	return items, nil
}
