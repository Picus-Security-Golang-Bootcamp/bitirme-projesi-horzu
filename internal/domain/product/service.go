package product

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type ProductService struct {
	repo Repository
}

type Service interface {
	GetAll(ctx context.Context, page, pageSize int) ([]Product,int64, error) 
	CreateProduct(ctx context.Context, name string, desc string, count int64, price float64, cid string) error
	DeleteProduct(ctx context.Context, sku string) error
	UpdateProduct(ctx context.Context, product *Product) error
	SearchProduct(ctx context.Context, text string, page, pageSize int) ([]Product,int64, error)
	UpdateProductQuantityForOrder(ctx context.Context,itemList []Product, amount []int64) error 
}

func NewProductService(repo *ProductRepository) Service {
	if repo == nil{
		return nil
	}
	return &ProductService{
		repo: repo,
	}

}

func (p *ProductService) GetAll(ctx context.Context, page, pageSize int) ([]Product,int64, error)  {
	products, count, err := p.repo.GetAll(ctx, page, pageSize)
	if err != nil {
		return nil, 0, err
	}


	return products, count, nil
}

func (p *ProductService) CreateProduct(ctx context.Context, name string, desc string, count int64, price float64, cid string) error {
	newProduct := NewProduct(name, desc, count, price, cid)
	_, err := p.repo.Create(ctx, newProduct)
	return err
}

func (c *ProductService) DeleteProduct(ctx context.Context, sku string) error {
	err := c.repo.Delete(ctx, sku)
	return err
}

func (c *ProductService) UpdateProduct(ctx context.Context, product *Product) error {
	changedProduct, err := c.repo.GetBySku(ctx, product.SKU)
	if err!=nil{
		return errors.New("record not found")
	}
	changedProduct.UpdateProduct(product.Name, product.SKU, product.Description, product.CategoryId, product.Stock, product.Price)

	_, err = c.repo.Update(ctx, changedProduct)

	return err
}

// SearchProduct finds Products that matches their sku number or names with given str field
func (c *ProductService) SearchProduct(ctx context.Context, text string, page, pageSize int) ([]Product,int64, error)  {
	products, count, err := c.repo.SearchByNameOrSku(ctx, text, page, pageSize)
	if err!=nil{
		return nil, 0 ,err
	}

	return products, count, nil
}

func (c *ProductService) UpdateProductQuantityForOrder(ctx context.Context,itemList []Product, amount []int64) error {

	for index, item := range itemList {
		product, err := c.repo.GetBySku(ctx, item.SKU)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		err1 := product.UpdateQuantity(amount[index])
		if err1 != nil {
			return err1
		}
	}

	for index, item := range itemList {
		product, _ := c.repo.GetBySku(ctx, item.SKU)
		product.UpdateQuantity(amount[index])
		c.repo.Update(ctx, product)
	}

	return nil
}