package service

import (
	"context"
	"time"

	"github.com/jimroxodezi/build-ms/internal/models"
	"github.com/jimroxodezi/build-ms/internal/repository"
)

type ProductService struct {
	productRepo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{productRepo: repo}
}

func (s *ProductService) CreateProduct(ctx context.Context, product models.Product) error {
	if err := product.Validate(); err != nil {
		return err
	}
	product.CreatedOn = time.Now()
	return s.productRepo.CreateProduct(product)
}

// 
func (s *ProductService) GetProduct(ctx context.Context, id int) (*models.Product, error) {
	return s.productRepo.FindProduct(id)
}

func (s *ProductService) GetProducts(ctx context.Context) ([]*models.Product, error) {
	return s.productRepo.FindProducts()
}

func (s *ProductService) UpdateProduct(ctx context.Context, id int, product models.Product) error {
	if err := product.Validate(); err != nil {
		return err
	}
	product.UpdatedOn = time.Now()
	return s.productRepo.UpdateProduct(id, &product)
}