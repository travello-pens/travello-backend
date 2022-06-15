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

func (r *RepositoryMysqlLayer) GetProductByID(id int) (product model.TravelAgentProduct, err error) {
	res := r.DB.Table("travel_agents ta").
		Joins("LEFT JOIN products p ON p.id_travel_agent = ta.id").
		Joins("LEFT JOIN locations l ON p.id_location = l.id").
		Select("ta.name_agent, ta.address, ta.phone_number, p.name_product, l.name_location, p.price, p.description, p.length_of_stay, p.address_of_product, p.tax, p.lodging, p.airport_pickup, p.refund").
		Where("ta.id = ?", id).
		Find(&product)

	// res := r.DB.Where("id = ?", id).Find(&product)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *RepositoryMysqlLayer) GetProductsByLocation(location string) (product []model.Product, err error) {
	var loc model.Location
	res1 := r.DB.Where("name_location = ?", location).Find(&loc)
	if res1.RowsAffected < 1 {
		err = fmt.Errorf("location not found")
	}

	res2 := r.DB.Where("id_location = ?", loc.ID).Find(&product)
	if res2.RowsAffected < 1 {
		err = fmt.Errorf("product not found")
	}

	return
}

func (r *RepositoryMysqlLayer) GetProductsByAgent(agent string) (product []model.Product, err error) {
	var agt model.TravelAgent
	res1 := r.DB.Where("name_agent = ?", agent).Find(&agt)
	if res1.RowsAffected < 1 {
		err = fmt.Errorf("travel agent not found")
	}

	res2 := r.DB.Where("id_travel_agent = ?", agt.ID).Find(&product)
	if res2.RowsAffected < 1 {
		err = fmt.Errorf("product not found")
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
