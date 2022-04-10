package productCategory

import (
	"github.com/horzu/golang/cart-api/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductCategoryRepository struct {
	db *gorm.DB
}

func NewProductCategoryRepository(db *gorm.DB) *ProductCategoryRepository {
	return &ProductCategoryRepository{db: db}
}

func (p *ProductCategoryRepository) create(b *models.ProductCategory) (*models.ProductCategory, error) {
	zap.L().Debug("productcategory.repo.create", zap.Reflect("productcategory", b))

	if err := p.db.Create(b).Error; err != nil {
		zap.L().Error("productcategory.repo.Create failed to create productcategory", zap.Error(err))
		return nil, err
	}

	return b, nil
}

func (p *ProductCategoryRepository) getAll() (*[]models.ProductCategory, error) {
	zap.L().Debug("productcategory.repo.getAll")

	var bs = &[]models.ProductCategory{}
	if err := p.db.Preload("productcategory").Find(&bs).Error; err != nil {
		zap.L().Error("productcategoryrepo.getAll failed to get productcategories", zap.Error(err))
		return nil, err
	}

	return bs, nil
}

func (p *ProductCategoryRepository) getByID(id string) (*models.ProductCategory, error) {
	zap.L().Debug("productcategory.repo.getByID", zap.Reflect("id", id))

	var category = &models.ProductCategory{}
	if result := p.db.First(&category, id); result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}

func (p *ProductCategoryRepository) update(a *models.ProductCategory) (*models.ProductCategory, error) {
	zap.L().Debug("productcategory.repo.update", zap.Reflect("productcategory", a))

	if result := p.db.Save(&a); result.Error != nil {
		return nil, result.Error
	}

	return a, nil
}

func (p *ProductCategoryRepository) delete(id string) error {
	zap.L().Debug("productcategory.repo.delete", zap.Reflect("id", id))

	category, err := p.getByID(id)
	if err != nil {
		return err
	}

	if result := p.db.Delete(&category); result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *ProductCategoryRepository) Migration() {
	p.db.AutoMigrate(&models.ProductCategory{})
}