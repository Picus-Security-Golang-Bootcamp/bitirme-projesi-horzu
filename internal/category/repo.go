package category

import (
	"github.com/horzu/golang/cart-api/internal/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (p *CategoryRepository) Migration() {
	p.db.AutoMigrate(&models.Category{})
}

func (p *CategoryRepository) create(b *models.Category) (*models.Category, error) {
	zap.L().Debug("category.repo.create", zap.Reflect("category", b))

	if err := p.db.Create(b).Error; err != nil {
		zap.L().Error("category.repo.Create failed to create category", zap.Error(err))
		return nil, err
	}

	return b, nil
}

func (p *CategoryRepository) getAll() (*[]models.Category, error) {
	zap.L().Debug("category.repo.getAll")

	var bs = &[]models.Category{}
	if err := p.db.Preload("category").Find(&bs).Error; err != nil {
		zap.L().Error("category.repo.getAll failed to get categories", zap.Error(err))
		return nil, err
	}

	return bs, nil
}

func (p *CategoryRepository) getByID(id string) (*models.Category, error) {
	zap.L().Debug("category.repo.getByID", zap.Reflect("id", id))

	var category = &models.Category{}
	if result := p.db.First(&category, id); result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}

func (p *CategoryRepository) update(a *models.Category) (*models.Category, error) {
	zap.L().Debug("category.repo.update", zap.Reflect("category", a))

	if result := p.db.Save(&a); result.Error != nil {
		return nil, result.Error
	}

	return a, nil
}

func (p *CategoryRepository) delete(id string) error {
	zap.L().Debug("category.repo.delete", zap.Reflect("id", id))

	category, err := p.getByID(id)
	if err != nil {
		return err
	}

	if result := p.db.Delete(&category); result.Error != nil {
		return result.Error
	}

	return nil
}
