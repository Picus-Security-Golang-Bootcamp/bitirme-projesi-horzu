package orderItem

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type OrderItemRepository struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) *OrderItemRepository {
	return &OrderItemRepository{db: db}
}



func (oir *OrderItemRepository) create(oi *OrderItem) (*OrderItem, error) {
	zap.L().Debug("product.repo.create", zap.Reflect("product", oi))

	if err := oir.db.Create(oi).Error; err != nil {
		zap.L().Error("product.repo.Create failed to create product", zap.Error(err))
		return nil, err
	}

	return oi, nil
}

func (oir *OrderItemRepository) update(oi *OrderItem) (*OrderItem, error) {
	zap.L().Debug("product.repo.update", zap.Reflect("product", oi))

	if result := oir.db.Save(&oi); result.Error != nil {
		return nil, result.Error
	}

	return oi, nil
}


