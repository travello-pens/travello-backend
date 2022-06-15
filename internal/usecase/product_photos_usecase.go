package usecase

import (
	"mime/multipart"

	"travel-agent-backend/config"
	"travel-agent-backend/internal/adapter"
	"travel-agent-backend/internal/model"
)

type serviceProductPhoto struct {
	c    config.Config
	repo adapter.AdapterProductPhotoRepository
}

func (s *serviceProductPhoto) CreateProductPhotoService(photo model.ProductPhoto, file *multipart.FileHeader) error {
	photo.File_Name = file.Filename

	return s.repo.CreateProductPhoto(photo)
}

func (s *serviceProductPhoto) GetAllProductPhotosService() []model.ProductPhoto {
	return s.repo.GetAllProductPhotos()
}

func (s *serviceProductPhoto) GetProductPhotoByIDService(id int) (model.ProductPhoto, error) {
	return s.repo.GetProductPhotoByID(id)
}

func (s *serviceProductPhoto) GetProductPhotoByProductService(product string) ([]model.ProductPhoto, error) {
	return s.repo.GetProductPhotoByProduct(product)
}

func (s *serviceProductPhoto) UpdateProductPhotoByIDService(id int, photo model.ProductPhoto, file *multipart.FileHeader) error {
	photo.File_Name = file.Filename

	return s.repo.UpdateProductPhotoByID(id, photo)
}

func (s *serviceProductPhoto) DeleteProductPhotoByIDService(id int) error {
	return s.repo.DeleteProductPhotoByID(id)
}

func NewServiceProductPhoto(repo adapter.AdapterProductPhotoRepository, c config.Config) adapter.AdapterProductPhotoService {
	return &serviceProductPhoto{
		repo: repo,
		c:    c,
	}
}
