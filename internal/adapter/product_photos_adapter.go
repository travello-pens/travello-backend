package adapter

import "travel-agent-backend/internal/model"

type AdapterProductPhotoRepository interface {
	CreateProductPhoto(photo model.ProductPhoto) error
	GetAllProductPhotos() []model.ProductPhoto
	GetProductPhotoByID(id int) (photo model.ProductPhoto, err error)
	UpdateProductPhotoByID(id int, photo model.ProductPhoto) error
	DeleteProductPhotoByID(id int) error
}

type AdapterProductPhotoService interface {
	CreateProductPhotoService(photo model.ProductPhoto) error
	GetAllProductPhotosService() []model.ProductPhoto
	GetProductPhotoByIDService(id int) (photo model.ProductPhoto, err error)
	UpdateProductPhotoByIDService(id int, photo model.ProductPhoto) error
	DeleteProductPhotoByIDService(id int) error
}
