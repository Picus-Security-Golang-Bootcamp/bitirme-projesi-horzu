package order

import (
	"github.com/horzu/golang/cart-api/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (o *OrderRepository) create(b *models.Order) (*models.Order, error) {
	zap.L().Debug("order.repo.create", zap.Reflect("order", b))

	if err := o.db.Create(b).Error; err != nil {
		zap.L().Error("order.repo.Create failed to create order", zap.Error(err))
		return nil, err
	}

	return b, nil
}

func (o *OrderRepository) getAll() (*[]models.Order, error) {
	zap.L().Debug("Order.repo.getAll")

	var bs = &[]models.Order{}
	if err := o.db.Preload("Product").Find(&bs).Error; err != nil {
		zap.L().Error("order.repo.getAll failed to get orders", zap.Error(err))
		return nil, err
	}

	return bs, nil
}

func (o *OrderRepository) getByID(id string) (*models.Order, error) {
	zap.L().Debug("Order.repo.getByID", zap.Reflect("id", id))

	var order = &models.Order{}
	if result := o.db.Preload("Product").First(&order, id); result.Error != nil {
		return nil, result.Error
	}
	return order, nil
}

func (o *OrderRepository) update(a *models.Order) (*models.Order, error) {
	zap.L().Debug("order.repo.update", zap.Reflect("order", a))

	if result := o.db.Save(&a); result.Error != nil {
		return nil, result.Error
	}

	return a, nil
}

func (o *OrderRepository) delete(id string) error {
	zap.L().Debug("order.repo.delete", zap.Reflect("id", id))

	order, err := o.getByID(id)
	if err != nil {
		return err
	}

	if result := o.db.Delete(&order); result.Error != nil {
		return result.Error
	}

	return nil
}

func (o *OrderRepository) Migration() {
	o.db.AutoMigrate(&models.Order{})
}