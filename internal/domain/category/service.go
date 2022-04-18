package category

import (
	"context"
	"mime/multipart"

	csvHelper "github.com/horzu/golang/cart-api/pkg/csv"
)

type categoryService struct {
	repo Repository
}

type Service interface {
	ListAll(ctx context.Context, page int, pageSize int) ([]Category,int64, error)
	CreateBulk(ctx context.Context, fileHeader *multipart.FileHeader) (int, error)
}

func NewCategoryService(repo *categoryRepository) Service {
	if repo == nil {
		return nil
	}

	return &categoryService{repo: repo}
}

// BulkCreate creates categories by uploaded csv files.
func (s *categoryService) CreateBulk(ctx context.Context, fileHeader *multipart.FileHeader) (int, error) {
	categories := make([]*Category, 0)
	data, err := csvHelper.ReadCSV(fileHeader)
	if err != nil {
		return 0, err
	}

	for _, readCategories := range data {
		categories = append(categories, NewCategory(readCategories[0], readCategories[1]))
	}

	count, err := s.repo.BulkCreate(ctx, categories)
	if err != nil {
		return count, err
	}
	return count, nil
}

// ListAll returns all active categories
func (s *categoryService) ListAll(ctx context.Context, page int, pageSize int) ([]Category,int64, error) {
	categories, count, err := s.repo.ListAll(ctx, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	return categories, count, nil
}
