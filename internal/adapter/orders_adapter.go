package adapter

import "travel-agent-backend/internal/model"

type AdapterOrderRepository interface {
	CreateOrder(order model.Order) error
	GetAllOrders() []model.Order
	GetOrderByID(id int) (order model.Order, err error)
	GetOrderByProduct(product string) (order []model.Order, err error)
	UpdateOrderByID(id int, order model.Order) error
	DeleteOrderByID(id int) error
}

type AdapterOrderService interface {
	CreateOrderService(order model.Order) error
	GetAllOrdersService() []model.Order
	GetOrderByIDService(id int) (order model.Order, err error)
	GetOrderByProductService(product string) (order []model.Order, err error)
	UpdateOrderByIDService(id int, order model.Order) error
	DeleteOrderByIDService(id int) error
}
