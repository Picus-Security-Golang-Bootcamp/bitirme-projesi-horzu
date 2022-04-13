package cartItem

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

type CartItemRepository struct {
	db *gorm.DB
}

func NewCartCartItemRepository(db *gorm.DB) *CartItemRepository {
	return &CartItemRepository{
		db: db,
	}
}

func (c *CartItemRepository) Migration() {
	err := c.db.AutoMigrate(&CartItem{})
	if err != nil {
		log.Print(err)
	}
}

func (c *CartItemRepository) Create(cartItem *CartItem) error {
	result := c.db.Create(cartItem)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *CartItemRepository) Update(cartItem *CartItem) error {
	result := c.db.Save(&cartItem)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (c *CartItemRepository) DeleteById(id string) error {
	result := c.db.Delete(&CartItem{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// FindByID returns Item with given productId and cartId
func (c *CartItemRepository) FindByID(productId string, cartId string) (*CartItem, error) {
	var item *CartItem

	err := c.db.Where(&CartItem{ProductId: productId, CartId: cartId}).First(&item).Error
	if err != nil {
		return nil, errors.New("cart item not found")
	}
	return item, nil
}

// GetItems return items in cart
func (c *CartItemRepository) GetItems(cartId string) ([]CartItem, error) {
	var cartItems []CartItem
	err := c.db.Where(&CartItem{CartId: cartId}).Find(&cartItems).Error
	if err != nil {
		return nil, err
	}
	for i, item := range cartItems {
		err := c.db.Model(item).Association("Product").Find(&cartItems[i].Product)
		if err != nil {
			return nil, err
		}
	}
	return cartItems, nil
}