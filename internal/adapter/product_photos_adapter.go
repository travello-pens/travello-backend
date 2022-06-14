package adapter

import (
	"mime/multipart"
	"travel-agent-backend/internal/model"
)

type AdapterProductPhotoRepository interface {
	CreateProductPhoto(photo model.ProductPhoto) error
	GetAllProductPhotos() []model.ProductPhoto
	GetProductPhotoByID(id int) (photo model.ProductPhoto, urlImage string, err error)
	GetProductPhotoByProduct(product string) (photo model.ProductPhoto, urlImage string, err error)
	UpdateProductPhotoByID(id int, photo model.ProductPhoto) error
	DeleteProductPhotoByID(id int) error
}

type AdapterProductPhotoService interface {
	CreateProductPhotoService(photo model.ProductPhoto, file *multipart.FileHeader) error
	GetAllProductPhotosService() []model.ProductPhoto
	GetProductPhotoByIDService(id int) (model.ProductPhoto, string, error)
	GetProductPhotoByProductService(product string) (model.ProductPhoto, string, error)
	UpdateProductPhotoByIDService(id int, photo model.ProductPhoto, file *multipart.FileHeader) error
	DeleteProductPhotoByIDService(id int) error
}
