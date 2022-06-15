package usecase

import (
	"travel-agent-backend/config"
	"travel-agent-backend/internal/adapter"
	"travel-agent-backend/internal/model"
)

type serviceOrder struct {
	c    config.Config
	repo adapter.AdapterOrderRepository
}

func (s *serviceOrder) CreateOrderService(order model.Order) error {
	return s.repo.CreateOrder(order)
}

func (s *serviceOrder) GetAllOrdersService() []model.Order {
	return s.repo.GetAllOrders()
}

func (s *serviceOrder) GetOrderByIDService(id int) (order model.Order, err error) {
	return s.repo.GetOrderByID(id)
}

func (s *serviceOrder) GetOrderByProductService(product string) (order []model.Order, err error) {
	return s.repo.GetOrderByProduct(product)
}

func (s *serviceOrder) UpdateOrderByIDService(id int, order model.Order) error {
	return s.repo.UpdateOrderByID(id, order)
}

func (s *serviceOrder) DeleteOrderByIDService(id int) error {
	return s.repo.DeleteOrderByID(id)
}

func NewServiceOrder(repo adapter.AdapterOrderRepository, c config.Config) adapter.AdapterOrderService {
	return &serviceOrder{
		repo: repo,
		c:    c,
	}
}
