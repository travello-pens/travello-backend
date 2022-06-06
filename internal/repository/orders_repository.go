package repository

import (
	"fmt"

	"travel-agent-backend/internal/adapter"
	"travel-agent-backend/internal/model"

	"gorm.io/gorm"
)

func (r *RepositoryMysqlLayer) CreateOrder(order model.Order) error {
	res := r.DB.Create(&order)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *RepositoryMysqlLayer) GetAllOrders() []model.Order {
	orders := []model.Order{}
	r.DB.Find(&orders)

	return orders
}

func (r *RepositoryMysqlLayer) GetOrderByID(id int) (order model.Order, err error) {
	res := r.DB.Where("id = ?", id).Find(&order)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *RepositoryMysqlLayer) UpdateOrderByID(id int, order model.Order) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&order)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *RepositoryMysqlLayer) DeleteOrderByID(id int) error {
	res := r.DB.Delete(&model.Order{ID: id})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlOrderRepository(db *gorm.DB) adapter.AdapterOrderRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
