package product

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, a *Product) (*Product, error)
	Update(ctx context.Context, a *Product) (*Product, error)
	Delete(ctx context.Context, sku string) error
	GetAll(ctx context.Context, page, pageSize int) ([]Product, int64, error)
	GetByID(ctx context.Context, id string) (*Product, error)
	GetBySku(ctx context.Context, sku string) (*Product, error)
	SearchByNameOrSku(ctx context.Context, str string, page, pageSize int) ([]*Product, int64, error) 
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) Migration() {
	p.db.AutoMigrate(&Product{})
}

func (p *ProductRepository) Create(ctx context.Context, a *Product) (*Product, error) {
	zap.L().Debug("product.repo.create", zap.Reflect("product", a))

	if err := p.db.Create(a).Error; err != nil {
		zap.L().Error("product.repo.Create failed to create product", zap.Error(err))
		return nil, err
	}

	return a, nil
}

func (p *ProductRepository) Update(ctx context.Context, a *Product) (*Product, error) {
	zap.L().Debug("product.repo.update", zap.Reflect("product", a))

	if result := p.db.Save(&a); result.Error != nil {
		return nil, result.Error
	}

	return a, nil
}

func (p *ProductRepository) Delete(ctx context.Context, sku string) error {
	zap.L().Debug("product.repo.delete", zap.Reflect("sku", sku))

	product, err := p.GetBySku(ctx, sku)
	if err != nil {
		return err
	}

	if result := p.db.Delete(&product); result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *ProductRepository) GetAll(ctx context.Context, page, pageSize int) ([]Product, int64, error) {
	zap.L().Debug("product.repo.getAll")

	var products []Product
	var count int64

	if err := p.db.Where("is_active = ?", true).Offset((page - 1) * pageSize).Limit(pageSize).Find(&products).Error; err != nil {
		zap.L().Error("product.repo.getAll failed to get products", zap.Error(err))
		return nil, 0, err
	}

	if err := p.db.Find(&products).Count(&count).Error; err != nil {
		zap.L().Error("category.repo.getAll failed to get categories count", zap.Error(err))
		return nil, 0, err
	}

	return products, count, nil
}

func (p *ProductRepository) GetByID(ctx context.Context, id string) (*Product, error) {
	zap.L().Debug("product.repo.getByID", zap.Reflect("id", id))

	var product = &Product{}
	if result := p.db.First(&product, id); result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (p *ProductRepository) GetBySku(ctx context.Context, sku string) (*Product, error) {
	zap.L().Debug("product.repo.GetBySku", zap.Reflect("sku", sku))

	var product *Product
	if result := p.db.Where("sku = ?", sku).First(&product); result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

// SearchByNameOrSku finds Products that matches their sku number or names with given str field
func (r *ProductRepository) SearchByNameOrSku(ctx context.Context, str string, page, pageSize int) ([]*Product, int64, error) {
	var products []*Product
	var count int64

	if result := r.db.Where("Name LIKE ? OR sku LIKE ?", "%" + str + "%", "%" + str + "%").Find(&products); result.Error != nil {
		return nil, 0, result.Error
	}

	if err := r.db.Where("Name LIKE ? OR sku LIKE ?", "%" + str + "%", "%" + str + "%").Find(&products).Count(&count).Error; err != nil {
		zap.L().Error("category.repo.SearchByNameOrSku failed to get products count", zap.Error(err))
		return nil, 0, err
	}

	fmt.Println(products)
	return products, count, nil
}

