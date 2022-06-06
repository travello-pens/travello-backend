package usecase

import (
	"travel-agent-backend/config"
	"travel-agent-backend/internal/adapter"
	"travel-agent-backend/internal/model"
)

type serviceProductPhoto struct {
	c    config.Config
	repo adapter.AdapterProductPhotoRepository
}

func (s *serviceProductPhoto) CreateProductPhotoService(photo model.ProductPhoto) error {
	return s.repo.CreateProductPhoto(photo)
}

func (s *serviceProductPhoto) GetAllProductPhotosService() []model.ProductPhoto {
	return s.repo.GetAllProductPhotos()
}

func (s *serviceProductPhoto) GetProductPhotoByIDService(id int) (photo model.ProductPhoto, err error) {
	return s.repo.GetProductPhotoByID(id)
}

func (s *serviceProductPhoto) UpdateProductPhotoByIDService(id int, photo model.ProductPhoto) error {
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
