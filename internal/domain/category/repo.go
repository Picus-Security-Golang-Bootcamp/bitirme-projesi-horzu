package category

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository interface {
	// Get returns the Category with the specified Category Id.
	Get(ctx context.Context, id string) (*Category, error)
	// Create saves a new Category in the storage.
	Create(ctx context.Context, Category *Category) error
	// BulkCreate saves a new Category list in the storage.
	BulkCreate(ctx context.Context, Category []*Category) (int, error)
	// Update updates the Category with given Is in the storage.
	Update(ctx context.Context, c *Category) (*Category, error)
	// Delete removes the Category with given Is from the storage.
	Delete(ctx context.Context, id string) error
	// ListAll returns all categories.
	ListAll(ctx context.Context, page int, pageSize int) ([]Category, error)
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

func (p *categoryRepository) Get(ctx context.Context, id string) (*Category, error) {
	zap.L().Debug("category.repo.getByID", zap.Reflect("id", id))

	var category = &Category{}
	if result := p.db.First(&category, id); result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}

func (p *categoryRepository) Create(ctx context.Context, b *Category) error {
	zap.L().Debug("category.repo.create", zap.Reflect("category", b))

	if err := p.db.WithContext(ctx).Create(b).Error; err != nil {
		zap.L().Error("category.repo.Create failed to create category", zap.Error(err))
		return err
	}

	return nil
}

func (p *categoryRepository) Update(ctx context.Context, c *Category) (*Category, error) {
	zap.L().Debug("category.repo.update", zap.Reflect("category", c))

	if result := p.db.Save(&c); result.Error != nil {
		return nil, result.Error
	}

	return c, nil
}

func (p *categoryRepository) Delete(ctx context.Context, id string) error {
	zap.L().Debug("category.repo.delete", zap.Reflect("id", id))

	category, err := p.Get(ctx, id)
	if err != nil {
		return err
	}

	if result := p.db.Delete(&category); result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *categoryRepository) ListAll(ctx context.Context, page int, pageSize int) ([]Category, error) {
	zap.L().Debug("category.repo.getAll")

	var bs []Category
	if err := p.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&bs).Error; err != nil {
		zap.L().Error("category.repo.getAll failed to get categories", zap.Error(err))
		return nil, err
	}

	return bs, nil
}

func (p *categoryRepository) BulkCreate(ctx context.Context, categories []*Category) (int, error) {
	zap.L().Debug("category.repo.BulkCreate", zap.Reflect("categories", categories))

	var count int64
	if result := p.db.Create(&categories).Count(&count); result.Error != nil {
		return 0, result.Error
	}
	return int(count), nil
}
