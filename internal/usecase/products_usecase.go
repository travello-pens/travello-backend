package usecase

import (
	"travel-agent-backend/config"
	"travel-agent-backend/internal/adapter"
	"travel-agent-backend/internal/model"
)

type serviceProduct struct {
	c    config.Config
	repo adapter.AdapterProductRepository
}

func (s *serviceProduct) CreateProductService(product model.Product) error {
	return s.repo.CreateProduct(product)
}

func (s *serviceProduct) GetAllProductsService() []model.Product {
	return s.repo.GetAllProducts()
}

func (s *serviceProduct) GetProductByIDService(id int) (product model.Product, err error) {
	return s.repo.GetProductByID(id)
}

func (s *serviceProduct) UpdateProductByIDService(id int, product model.Product) error {
	return s.repo.UpdateProductByID(id, product)
}

func (s *serviceProduct) DeleteProductByIDService(id int) error {
	return s.repo.DeleteProductByID(id)
}

func NewServiceProduct(repo adapter.AdapterProductRepository, c config.Config) adapter.AdapterProductService {
	return &serviceProduct{
		repo: repo,
		c:    c,
	}
}
