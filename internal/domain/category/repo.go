package category

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository interface {
	// BulkCreate saves a new Category list in the storage.
	BulkCreate(ctx context.Context, Category []*Category) (int, error)
	// ListAll returns all categories.
	ListAll(ctx context.Context, page int, pageSize int) ([]Category, int64, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

func (p *categoryRepository) Migration() {
	p.db.AutoMigrate(&Category{})
}

func (p *categoryRepository) ListAll(ctx context.Context, page int, pageSize int) ([]Category, int64, error) {
	zap.L().Debug("category.repo.getAll")

	var categories []Category
	var count int64

	if err := p.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&categories).Error; err != nil {
		zap.L().Error("category.repo.getAll failed to get categories", zap.Error(err))
		return nil, 0, err
	}
	if err := p.db.Find(&categories).Count(&count).Error; err != nil {
		zap.L().Error("category.repo.getAll failed to get categories count", zap.Error(err))
		return nil, 0, err
	}

	return categories, count, nil
}

func (p *categoryRepository) BulkCreate(ctx context.Context, categories []*Category) (int, error) {
	zap.L().Debug("category.repo.BulkCreate", zap.Reflect("categories", categories))

	var count int64
	if result := p.db.Create(&categories).Count(&count); result.Error != nil {
		return 0, result.Error
	}
	return int(count), nil
}
