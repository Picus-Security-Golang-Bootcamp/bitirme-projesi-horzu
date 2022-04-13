package product

import (
	"context"

	"github.com/horzu/golang/cart-api/pkg/pagination"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, a *Product) (*Product, error)
	Update(ctx context.Context, a *Product) (*Product, error)
	Delete(ctx context.Context, sku string) error
	GetAll(ctx context.Context, page *pagination.Pages) ([]Product, int64, error)
	GetByID(ctx context.Context, id string) (*Product, error)
	GetBySku(ctx context.Context, sku string) (*Product, error)
	SearchByNameOrSku(ctx context.Context, str string, page *pagination.Pages) ([]*Product, int, error)
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

func (p *ProductRepository) GetAll(ctx context.Context, page *pagination.Pages) ([]Product, int64, error) {
	zap.L().Debug("product.repo.getAll")

	var products []Product
	var pages int64

	if err := p.db.Where("is_active = ?", true).Offset((page.Page - 1) * page.PageSize).Limit(page.PageSize).Find(&products).Count(&pages).Error; err != nil {
		zap.L().Error("product.repo.getAll failed to get products", zap.Error(err))
		return nil, 0, err
	}

	return products, pages, nil
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
	if result := p.db.First(&product, "sku",sku); result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

// SearchByNameOrSku finds Products that matches their sku number or names with given str field
func (r *ProductRepository) SearchByNameOrSku(ctx context.Context, str string, page *pagination.Pages) ([]*Product, int, error) {
	var products []*Product
	convertedStr := "%" + str + "%"
	var count int64
	if result := r.db.Where("Name LIKE ? OR sku Like ?", convertedStr, convertedStr).Offset((page.Page - 1) * page.PageSize).Limit(page.PageSize).Find(&products).Count(&count); result.Error != nil {
		return nil, 0, result.Error
	}

	return products, int(count), nil
}
