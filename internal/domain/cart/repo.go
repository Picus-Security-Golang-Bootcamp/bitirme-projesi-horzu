package cart

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Repository encapsulates the logic to access basket from the data source.
type Repository interface {
	Get(ctx context.Context, id string) *Cart
	Create(ctx context.Context, b *Cart) error
	Update(ctx context.Context, a *Cart) (*Cart, error)
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]*Cart, error)
	GetByID(ctx context.Context, id string) (*Cart, error)
	FindOrCreateByUserID(ctx context.Context, userId string) (*Cart, error)
	FindByUserId(ctx context.Context, userId string) (*Cart, error)
}

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (o *CartRepository) Migration() {
	o.db.AutoMigrate(&Cart{})
}

func (o *CartRepository) Get(ctx context.Context, id string) *Cart {
	var cart *Cart
	o.db.WithContext(ctx).Where("Id = ?", id).Find(&cart)

	return cart
}

func (o *CartRepository) Create(ctx context.Context, c *Cart) error {
	zap.L().Debug("cart.repo.create", zap.Reflect("cart", c))

	if err := o.db.Create(c).Error; err != nil {
		zap.L().Error("cart.repo.Create failed to create cart", zap.Error(err))
		return err
	}

	return nil
}

func (o *CartRepository) Update(ctx context.Context, c *Cart) (*Cart, error) {
	zap.L().Debug("cart.repo.update", zap.Reflect("cart", c))

	if result := o.db.Save(&c); result.Error != nil {
		return nil, result.Error
	}

	return c, nil
}

func (o *CartRepository) Delete(ctx context.Context, id string) error {
	zap.L().Debug("cart.repo.delete", zap.Reflect("id", id))

	cart, err := o.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if result := o.db.Delete(&cart); result.Error != nil {
		return result.Error
	}

	return nil
}

// getAll returns all carts infos
func (o *CartRepository) GetAll(ctx context.Context) ([]*Cart, error) {
	zap.L().Debug("cart.repo.getAll")

	var bs []*Cart
	if err := o.db.Find(&bs).Error; err != nil {
		zap.L().Error("cart.repo.getAll failed to get carts", zap.Error(err))
		return nil, err
	}

	return bs, nil
}

// getByID returns cart info of the user by id
func (o *CartRepository) GetByID(ctx context.Context, id string) (*Cart, error) {
	zap.L().Debug("Order.repo.getByID", zap.Reflect("id", id))

	var cart = &Cart{}
	if result := o.db.First(&cart, id); result.Error != nil {
		return nil, result.Error
	}
	return cart, nil
}

// FindOrCreateByUserID returns the cart of the user if exists
// If cart does not exist, it creates a new one
func (o *CartRepository) FindOrCreateByUserID(ctx context.Context, userId string) (*Cart, error) {
	var cart *Cart
	err := o.db.Where(Cart{UserID: userId}).Attrs(NewCart(userId)).FirstOrCreate(&cart).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}

func (o *CartRepository) FindByUserId(ctx context.Context, userId string) (*Cart, error) {
	var cart *Cart

	result := o.db.Where("user_id = ?", userId).Limit(1).Find(&cart)

	if result.Error != nil {
		return nil, result.Error
	}

	return cart, nil
}
