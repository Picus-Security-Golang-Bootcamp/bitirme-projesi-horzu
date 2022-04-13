package category

import (
	"context"
	"errors"
	"mime/multipart"

	csvHelper "github.com/horzu/golang/cart-api/pkg/csv"
	"github.com/horzu/golang/cart-api/pkg/pagination"
)

type categoryService struct {
	repo Repository
}

type Service interface {
	ListAll(ctx context.Context, page *pagination.Pages) ([]Category, error)
	Create(ctx context.Context, category *Category) error
	CreateBulk(ctx context.Context, fileHeader *multipart.FileHeader) (int, error)
}

func NewCategoryService(repo *categoryRepository) Service {
	if repo == nil {
		return nil
	}

	return &categoryService{repo: repo}
}

func (s *categoryService) Create(ctx context.Context, category *Category) error {
	existCity, _ := s.repo.Get(ctx, category.Id)
	if existCity != nil {
		return errors.New("It's already exist")
	}

	err := s.repo.Create(ctx, category)
	if err != nil {
		return err
	}

	return nil
}

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

func (s *categoryService) ListAll(ctx context.Context, page *pagination.Pages) ([]Category, error) {
	var categories []Category

	categories, err := s.repo.ListAll(ctx, page.Page, page.PageSize)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
