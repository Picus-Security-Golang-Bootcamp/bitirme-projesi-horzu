package orderProduct

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

func (p *ProductRepository) create(b *models.Product) (*models.Product, error) {
	zap.L().Debug("product.repo.create", zap.Reflect("product", b))

	if err := p.db.Create(b).Error; err != nil {
		zap.L().Error("product.repo.Create failed to create product", zap.Error(err))
		return nil, err
	}

	return b, nil
}

func (p *ProductRepository) getAll() (*[]models.Product, error) {
	zap.L().Debug("product.repo.getAll")

	var bs = &[]models.Product{}
	if err := p.db.Preload("Product").Find(&bs).Error; err != nil {
		zap.L().Error("product.repo.getAll failed to get products", zap.Error(err))
		return nil, err
	}

	return bs, nil
}

func (p *ProductRepository) getByID(id string) (*models.Product, error) {
	zap.L().Debug("product.repo.getByID", zap.Reflect("id", id))

	var product = &models.Product{}
	if result := p.db.First(&product, id); result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (p *ProductRepository) update(a *models.Product) (*models.Product, error) {
	zap.L().Debug("product.repo.update", zap.Reflect("product", a))

	if result := p.db.Save(&a); result.Error != nil {
		return nil, result.Error
	}

	return a, nil
}

func (p *ProductRepository) delete(id string) error {
	zap.L().Debug("product.repo.delete", zap.Reflect("id", id))

	product, err := p.getByID(id)
	if err != nil {
		return err
	}

	if result := p.db.Delete(&product); result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *ProductRepository) Migration() {
	p.db.AutoMigrate(&models.Product{})
}