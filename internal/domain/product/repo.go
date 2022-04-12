package product

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) Migration() {
	p.db.AutoMigrate(&Product{})
}

func (p *ProductRepository) Create(b *Product) (*Product, error) {
	zap.L().Debug("product.repo.create", zap.Reflect("product", b))

	if err := p.db.Create(b).Error; err != nil {
		zap.L().Error("product.repo.Create failed to create product", zap.Error(err))
		return nil, err
	}

	return b, nil
}

func (p *ProductRepository) Update(a *Product) (*Product, error) {
	zap.L().Debug("product.repo.update", zap.Reflect("product", a))

	if result := p.db.Save(&a); result.Error != nil {
		return nil, result.Error
	}

	return a, nil
}

func (p *ProductRepository) Delete(id string) error {
	zap.L().Debug("product.repo.delete", zap.Reflect("id", id))

	product, err := p.GetByID(id)
	if err != nil {
		return err
	}

	if result := p.db.Delete(&product); result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *ProductRepository) GetAll(pageIndex, pageSize int) ([]Product, int64, error) {
	zap.L().Debug("product.repo.getAll")

	var products []Product
	var pages int64

	if err := p.db.Where("IsDeleted = ?", 0).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&pages).Error; err != nil {
		zap.L().Error("product.repo.getAll failed to get products", zap.Error(err))
		return nil, 0, err
	}

	return products, pages, nil
}

func (p *ProductRepository) GetByID(id string) (*Product, error) {
	zap.L().Debug("product.repo.getByID", zap.Reflect("id", id))

	var product = &Product{}
	if result := p.db.First(&product, id); result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (p *ProductRepository) GetBySku(sku string) (*Product, error) {
	zap.L().Debug("product.repo.GetBySku", zap.Reflect("sku", sku))

	var product *Product
	if result := p.db.First(&product, sku); result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

// SearchByNameOrSku finds Products that matches their sku number or names with given str field
func (r *ProductRepository) SearchByNameOrSku(str string, pageIndex, pageSize int) ([]*Product, int) {
	var products []*Product
	convertedStr := "%" + str + "%"
	var count int64
	r.db.Where(
		"Name LIKE ? OR SKU Like ?", convertedStr,
		convertedStr).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)

	return products, int(count)
}