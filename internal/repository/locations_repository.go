package repository

import (
	"fmt"

	"travel-agent-backend/internal/adapter"
	"travel-agent-backend/internal/model"

	"gorm.io/gorm"
)

func (r *RepositoryMysqlLayer) CreateLocation(location model.Location) error {
	res := r.DB.Create(&location)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *RepositoryMysqlLayer) GetAllLocations() []model.Location {
	locations := []model.Location{}
	r.DB.Find(&locations)

	return locations
}

func (r *RepositoryMysqlLayer) GetLocationByID(id int) (location model.Location, err error) {
	res := r.DB.Where("id = ?", id).Find(&location)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *RepositoryMysqlLayer) UpdateLocationByID(id int, location model.Location) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&location)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *RepositoryMysqlLayer) DeleteLocationByID(id int) error {
	res := r.DB.Delete(&model.Location{ID: id})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlLocationRepository(db *gorm.DB) adapter.AdapterLocationRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
