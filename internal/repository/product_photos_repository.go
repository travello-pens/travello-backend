package repository

import (
	"fmt"

	"travel-agent-backend/internal/adapter"
	"travel-agent-backend/internal/model"

	"gorm.io/gorm"
)

func (r *RepositoryMysqlLayer) CreateProductPhoto(photo model.ProductPhoto) error {
	res := r.DB.Create(&photo)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *RepositoryMysqlLayer) GetAllProductPhotos() []model.ProductPhoto {
	photos := []model.ProductPhoto{}
	r.DB.Find(&photos)

	return photos
}

func (r *RepositoryMysqlLayer) GetProductPhotoByID(id int) (photo model.ProductPhoto, err error) {
	res := r.DB.Where("id = ?", id).Find(&photo)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *RepositoryMysqlLayer) UpdateProductPhotoByID(id int, photo model.ProductPhoto) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&photo)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *RepositoryMysqlLayer) DeleteProductPhotoByID(id int) error {
	res := r.DB.Delete(&model.ProductPhoto{ID: id})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlProductPhotoRepository(db *gorm.DB) adapter.AdapterProductPhotoRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
