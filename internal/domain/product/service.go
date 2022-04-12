package product

import (
	"github.com/horzu/golang/cart-api/pkg/pagination"
)

type ProductService struct {
	productRepository ProductRepository
}

func NewProductService(productRepository ProductRepository) *ProductService {
	productRepository.Migration()
	return &ProductService{
		productRepository: productRepository,
	}

}

func (p *ProductService) GetAll(page *pagination.Pages) *pagination.Pages {
	products, count, err := p.productRepository.GetAll(page.Page, page.PageSize)
	if err!=nil{
		return nil
	}
	page.Items = products
	page.TotalCount = int(count)
	return page
}

func (p *ProductService) CreateProduct(name string, desc string, count uint, price float64, cid uint) error {
	newProduct := NewProduct(name, desc, count, price, cid)
	_, err := p.productRepository.Create(newProduct)
	return err
}

func (c *ProductService) DeleteProduct(sku string) error {
	err := c.productRepository.Delete(sku)
	return err
}

func (c *ProductService) UpdateProduct(product *Product) error {
	_, err := c.productRepository.Update(product)
	return err
}

// SearchProduct finds Products that matches their sku number or names with given str field
func (c *ProductService) SearchProduct(text string, page *pagination.Pages) *pagination.Pages {
	products, count := c.productRepository.SearchByNameOrSku(text, page.Page, page.PageSize)
	page.Items = products
	page.TotalCount = count
	return page
}