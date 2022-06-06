package repository

import (
	"fmt"

	"travel-agent-backend/internal/adapter"
	"travel-agent-backend/internal/model"

	"gorm.io/gorm"
)

func (r *RepositoryMysqlLayer) CreateProduct(product model.Product) error {
	res := r.DB.Create(&product)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *RepositoryMysqlLayer) GetAllProducts() []model.Product {
	products := []model.Product{}
	r.DB.Find(&products)

	return products
}

func (r *RepositoryMysqlLayer) GetProductByID(id int) (product model.Product, err error) {
	res := r.DB.Where("id = ?", id).Find(&product)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *RepositoryMysqlLayer) UpdateProductByID(id int, product model.Product) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&product)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *RepositoryMysqlLayer) DeleteProductByID(id int) error {
	res := r.DB.Delete(&model.Product{ID: id})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlProductRepository(db *gorm.DB) adapter.AdapterProductRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
