package usecase

import (
	"travel-agent-backend/config"
	"travel-agent-backend/internal/adapter"
	"travel-agent-backend/internal/model"
)

type serviceLocation struct {
	c    config.Config
	repo adapter.AdapterLocationRepository
}

func (s *serviceLocation) CreateLocationService(location model.Location) error {
	return s.repo.CreateLocation(location)
}

func (s *serviceLocation) GetAllLocationsService() []model.Location {
	return s.repo.GetAllLocations()
}

func (s *serviceLocation) GetLocationByIDService(id int) (location model.Location, err error) {
	return s.repo.GetLocationByID(id)
}

func (s *serviceLocation) UpdateLocationByIDService(id int, location model.Location) error {
	return s.repo.UpdateLocationByID(id, location)
}

func (s *serviceLocation) DeleteLocationByIDService(id int) error {
	return s.repo.DeleteLocationByID(id)
}

func NewServiceLocation(repo adapter.AdapterLocationRepository, c config.Config) adapter.AdapterLocationService {
	return &serviceLocation{
		repo: repo,
		c:    c,
	}
}
