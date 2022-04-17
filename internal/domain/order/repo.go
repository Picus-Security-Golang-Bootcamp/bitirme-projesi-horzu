package order

import (
	"context"

	"github.com/horzu/golang/cart-api/internal/domain/order/orderItem"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, o *Order) error
	DeleteById(ctx context.Context, id string) error
	GetAllByUser(ctx context.Context, id string) ([]*Order, error)
	GetByID(ctx context.Context, id string) (*Order, error)
	CreateWithCartItems(ctx context.Context, items []*orderItem.OrderItem) error 
}

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (or *OrderRepository) Migration() {
	or.db.AutoMigrate(&Order{})
}

func (or *OrderRepository) Create(ctx context.Context, o *Order) error {
	zap.L().Debug("order.repo.create", zap.Reflect("order", o))

	if err := or.db.Create(o).Error; err != nil {
		zap.L().Error("order.repo.Create failed to create order", zap.Error(err))
		return err
	}

	return nil
}
func (or *OrderRepository) CreateWithCartItems(ctx context.Context, items []*orderItem.OrderItem) error {
	zap.L().Debug("order.repo.create failed to create order", zap.Reflect("order", items))

	if err := or.db.Create(items).Error; err != nil {
		zap.L().Error("order.repo.Create failed to create order", zap.Error(err))
		return err
	}

	return nil
}

func (or *OrderRepository) DeleteById(ctx context.Context, id string) error {
	zap.L().Debug("order.repo.delete", zap.Reflect("id", id))

	if result := or.db.Where("id = ?", id).Delete(&Order{}); result.Error != nil {
		return result.Error
	}

	return nil
}

func (or *OrderRepository) GetAllByUser(ctx context.Context, id string) ([]*Order, error) {
	zap.L().Debug("Order.repo.getAll")

	var orders []*Order
	if err := or.db.Where("is_canceled = ?", false).Preload("Items").Where("user_id = ?", id).Find(&orders).Error; err != nil {
		zap.L().Error("order.repo.getAll failed to get all orders", zap.Error(err))
		return nil, err
	}

	return orders, nil
}

func (o *OrderRepository) GetByID(ctx context.Context, id string) (*Order, error) {
	zap.L().Debug("Order.repo.getByID", zap.Reflect("id", id))

	var order *Order
	if result := o.db.Preload("Items").Preload("Items.Product").Where("id = ?", id).First(&order); result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}
