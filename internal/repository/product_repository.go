package repository

import (
	"time"

	"github.com/jimroxodezi/build-ms/internal/models"
)

type ProductRepository interface {
	CreateProduct(product models.Product) error
	UpdateProduct(id int, p *models.Product) error
	DeleteProduct(id int) error
	FindProduct(id int) (*models.Product, error)
	FindProducts() ([]*models.Product, error)
}


type InMemoryProductRepository struct {
	productStore map[int]*models.Product
}

func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		productStore: make(map[int]*models.Product),
	}
}

func (r *InMemoryProductRepository) CreateProduct(product models.Product) error {
	product.ID = len(r.productStore) + 1
	product.CreatedOn = time.Now()
	product.UpdatedOn = time.Now()
	r.productStore[product.ID] = &product
	return nil
}

func (r *InMemoryProductRepository) DeleteProduct(id int) error {
	_, ok := r.productStore[id]
	if !ok {
		return models.ErrProductNotFound
	}
	delete(r.productStore, id)
	return nil
}

func (r *InMemoryProductRepository) UpdateProduct(id int, p *models.Product) error {
	prod, err := r.findProduct(id)
	if err != nil {
		return err
	}

	prod.ID = id
	*prod = *p
	prod.UpdatedOn = time.Now()
	return nil
}

// FindProduct finds a product by its ID and returns it.
func (r *InMemoryProductRepository) FindProduct(id int) (*models.Product, error) {
	return r.findProduct(id)
}

func (r *InMemoryProductRepository) FindProducts() ([]*models.Product, error) {
	products := make([]*models.Product, 0, len(r.productStore))
	for _, p := range r.productStore {
		products = append(products, p)
	}
	return products, nil
}

func (r *InMemoryProductRepository) findProduct(id int) (*models.Product, error) {
	p, ok := r.productStore[id]
	if !ok {
		return nil, models.ErrProductNotFound
	}

	return p, nil
}