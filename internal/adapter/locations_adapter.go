package adapter

import "travel-agent-backend/internal/model"

type AdapterLocationRepository interface {
	CreateLocation(location model.Location) error
	GetAllLocations() []model.Location
	GetLocationByID(id int) (location model.Location, err error)
	UpdateLocationByID(id int, location model.Location) error
	DeleteLocationByID(id int) error
}

type AdapterLocationService interface {
	CreateLocationService(location model.Location) error
	GetAllLocationsService() []model.Location
	GetLocationByIDService(id int) (location model.Location, err error)
	UpdateLocationByIDService(id int, location model.Location) error
	DeleteLocationByIDService(id int) error
}
