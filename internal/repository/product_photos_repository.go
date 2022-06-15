package repository

import (
	"fmt"
	"log"
	"os"

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

func (r *RepositoryMysqlLayer) GetProductPhotoByID(id int) (model.ProductPhoto, error) {
	var photo model.ProductPhoto

	res := r.DB.Where("id = ?", id).Find(&photo)
	if res.RowsAffected < 1 {
		return photo, fmt.Errorf("not found")
	}

	// filePath := "storage/"
	// urlImage := filePath + photo.File_Name

	return photo, nil
}

func (r *RepositoryMysqlLayer) GetProductPhotoByProduct(product string) ([]model.ProductPhoto, error) {
	var prod model.Product
	res1 := r.DB.Where("name_product = ?", product).Find(&prod)
	if res1.RowsAffected < 1 {
		return []model.ProductPhoto{}, fmt.Errorf("not found")
	}

	var photo []model.ProductPhoto
	res2 := r.DB.Where("id_product = ?", prod.ID).Find(&photo)
	if res2.RowsAffected < 1 {
		return photo, fmt.Errorf("order not found")
	}

	// filePath := "storage/"
	// urlImage := filePath + photo.File_Name

	return photo, nil
}

func (r *RepositoryMysqlLayer) UpdateProductPhotoByID(id int, photo model.ProductPhoto) error {
	// Finding data
	var img model.ProductPhoto
	res1 := r.DB.Where("id = ?", id).Find(&img)
	if res1.RowsAffected < 1 {
		return fmt.Errorf("error find image")
	}

	// Delete file in local storage
	filePath := "storage/"
	err := os.Remove(filePath + img.File_Name)
	if err != nil {
		log.Fatal(err)
	}

	// Update data in database
	res := r.DB.Where("id = ?", id).UpdateColumns(&photo)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *RepositoryMysqlLayer) DeleteProductPhotoByID(id int) error {
	// Finding data
	var img model.ProductPhoto
	res1 := r.DB.Where("id = ?", id).Find(&img)
	if res1.RowsAffected < 1 {
		return fmt.Errorf("error find image")
	}

	// Delete file in local storage
	filePath := "storage/"
	err := os.Remove(filePath + img.File_Name)
	if err != nil {
		log.Fatal(err)
	}

	// Delete data in database
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
