package orderItem

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/horzu/golang/cart-api/internal/domain/cart/cartItem"
	"github.com/horzu/golang/cart-api/internal/domain/product"
	"gorm.io/gorm"
)

type OrderItem struct {
	Id         string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	OrderId    string
	Quantity   uint
	
	ProductId  string
	Product *product.Product
}

func (u *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()

	return nil
}

func NewOrderItem(cartItem cartItem.CartItem) *OrderItem {
	item := &OrderItem{
		ProductId: cartItem.ProductId,
		Quantity:  cartItem.Quantity,
	}
	return item
}

// BeforeSave updates products' stock count after completing order
func (orderedItem *OrderItem) BeforeSave(tx *gorm.DB) (err error) {

	var currentProduct product.Product
	var currentOrderedItem OrderItem
	if err := tx.Where("ID = ?", orderedItem.ProductId).First(&currentProduct).Error; err != nil {
		return err
	}
	reservedStockCount := 0
	if err := tx.Where("ID = ?", orderedItem.Id).First(&currentOrderedItem).Error; err == nil {
		reservedStockCount = int(currentOrderedItem.Quantity)
	}
	newStockCount := int(currentProduct.Stock) + reservedStockCount - int(orderedItem.Quantity)
	if newStockCount < 0 {
		return errors.New("Not Enough Stock")
	}
	if err := tx.Model(&currentProduct).Update("stock", newStockCount).Error; err != nil {
		return err
	}
	if orderedItem.Quantity == 0 {
		err := tx.Unscoped().Delete(currentOrderedItem).Error
		return err
	}
	return
}