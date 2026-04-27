package service

import (
	"context"
	"time"

	"github.com/jimroxodezi/build-ms/models"
	"github.com/jimroxodezi/build-ms/repository"
)

// type ProductService interface {
// 	CreateProduct(ctx context.Context, product repository.ProductRepository) error
// 	GetProducts(ctx context.Context) ([]*repository.ProductRepository, error)
// }

type ProductService struct {
	productRepo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{productRepo: repo}
}

func (s *ProductService) CreateProduct(ctx context.Context, product models.Product) error {
	product.CreatedOn = time.Now()
	if err := product.Validate(); err != nil {
		return err
	}
	return s.productRepo.CreateProduct(product)
}

func (s *ProductService) GetProducts(ctx context.Context) ([]*models.Product, error) {
	return s.productRepo.FindAllProducts()
}

func (s *ProductService) UpdateProduct(ctx context.Context, id int, product models.Product) error {
	return s.productRepo.UpdateProduct(id, &product)
}