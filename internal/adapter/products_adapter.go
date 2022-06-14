package adapter

import "travel-agent-backend/internal/model"

type AdapterProductRepository interface {
	CreateProduct(product model.Product) error
	GetAllProducts() []model.Product
	GetProductByID(id int) (product model.Product, err error)
	GetProductsByLocation(location string) (product model.Product, err error)
	GetProductsByAgent(agent string) (product model.Product, err error)
	UpdateProductByID(id int, product model.Product) error
	DeleteProductByID(id int) error
}

type AdapterProductService interface {
	CreateProductService(product model.Product) error
	GetAllProductsService() []model.Product
	GetProductByIDService(id int) (product model.Product, err error)
	GetProductByLocationService(location string) (product model.Product, err error)
	GetProductByAgentService(agent string) (product model.Product, err error)
	UpdateProductByIDService(id int, product model.Product) error
	DeleteProductByIDService(id int) error
}
