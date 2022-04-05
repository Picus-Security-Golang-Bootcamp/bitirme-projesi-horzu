package product

import (
	"github.com/horzu/golang/cart-api/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (o *ProductRepository) create(b *models.Product) (*models.Product, error) {
	zap.L().Debug("product.repo.create", zap.Reflect("product", b))

	if err := o.db.Create(b).Error; err != nil {
		zap.L().Error("product.repo.Create failed to create product", zap.Error(err))
		return nil, err
	}

	return b, nil
}

func (o *ProductRepository) getAll() (*[]models.Product, error) {
	zap.L().Debug("product.repo.getAll")

	var bs = &[]models.Product{}
	if err := o.db.Preload("Product").Find(&bs).Error; err != nil {
		zap.L().Error("product.repo.getAll failed to get products", zap.Error(err))
		return nil, err
	}

	return bs, nil
}

func (o *ProductRepository) getByID(id string) (*models.Product, error) {
	zap.L().Debug("product.repo.getByID", zap.Reflect("id", id))

	var product = &models.Product{}
	if result := o.db.Preload("Category").First(&product, id); result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (o *ProductRepository) update(a *models.Product) (*models.Product, error) {
	zap.L().Debug("product.repo.update", zap.Reflect("product", a))

	if result := o.db.Save(&a); result.Error != nil {
		return nil, result.Error
	}

	return a, nil
}

func (o *ProductRepository) delete(id string) error {
	zap.L().Debug("product.repo.delete", zap.Reflect("id", id))

	product, err := o.getByID(id)
	if err != nil {
		return err
	}

	if result := o.db.Delete(&product); result.Error != nil {
		return result.Error
	}

	return nil
}

func (o *ProductRepository) Migration() {
	o.db.AutoMigrate(&models.Product{})
}