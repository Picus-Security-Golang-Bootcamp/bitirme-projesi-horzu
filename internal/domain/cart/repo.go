package cart

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (o *CartRepository) Migration() {
	o.db.AutoMigrate(&Cart{})
}

func (o *CartRepository) create(b *Cart) (*Cart, error) {
	zap.L().Debug("cart.repo.create", zap.Reflect("cart", b))

	if err := o.db.Create(b).Error; err != nil {
		zap.L().Error("cart.repo.Create failed to create cart", zap.Error(err))
		return nil, err
	}

	return b, nil
}

func (o *CartRepository) update(a *Cart) (*Cart, error) {
	zap.L().Debug("cart.repo.update", zap.Reflect("cart", a))

	if result := o.db.Save(&a); result.Error != nil {
		return nil, result.Error
	}

	return a, nil
}

func (o *CartRepository) delete(id string) error {
	zap.L().Debug("cart.repo.delete", zap.Reflect("id", id))

	cart, err := o.getByID(id)
	if err != nil {
		return err
	}

	if result := o.db.Delete(&cart); result.Error != nil {
		return result.Error
	}

	return nil
}

// getAll returns all carts infos
func (o *CartRepository) getAll() (*[]Cart, error) {
	zap.L().Debug("cart.repo.getAll")

	var bs = &[]Cart{}
	if err := o.db.Find(&bs).Error; err != nil {
		zap.L().Error("cart.repo.getAll failed to get carts", zap.Error(err))
		return nil, err
	}

	return bs, nil
}

// getByID returns cart info of the user by id
func (o *CartRepository) getByID(id string) (*Cart, error) {
	zap.L().Debug("Order.repo.getByID", zap.Reflect("id", id))

	var cart = &Cart{}
	if result := o.db.First(&cart, id); result.Error != nil {
		return nil, result.Error
	}
	return cart, nil
}

// FindOrCreateByUserID returns the cart of the user if exists
// If cart does not exist, it creates a new one
func (o *CartRepository) FindOrCreateByUserID(userId string) (*Cart, error) {
	var cart *Cart
	err := o.db.Where(Cart{UserID: userId}).Attrs(NewCart(userId)).FirstOrCreate(&cart).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}
