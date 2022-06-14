package adapter

import "travel-agent-backend/internal/model"

type AdapterTravelAgentSocialMediaRepository interface {
	CreateTravelAgentSocialMedia(media model.TravelAgentSocialMedia) error
	GetAllTravelAgentSocialMedias() []model.TravelAgentSocialMedia
	GetTravelAgentSocialMediaByID(id int) (media model.TravelAgentSocialMedia, err error)
	GetTravelAgentSocialMediaByTravelAgent(agent string) (media model.TravelAgentSocialMedia, err error)
	UpdateTravelAgentSocialMediaByID(id int, media model.TravelAgentSocialMedia) error
	DeleteTravelAgentSocialMediaByID(id int) error
}

type AdapterTravelAgentSocialMediaService interface {
	CreateTravelAgentSocialMediaService(media model.TravelAgentSocialMedia) error
	GetAllTravelAgentSocialMediasService() []model.TravelAgentSocialMedia
	GetTravelAgentSocialMediaByIDService(id int) (media model.TravelAgentSocialMedia, err error)
	GetTravelAgentSocialMediaByTravelAgentService(agent string) (media model.TravelAgentSocialMedia, err error)
	UpdateTravelAgentSocialMediaByIDService(id int, media model.TravelAgentSocialMedia) error
	DeleteTravelAgentSocialMediaByIDService(id int) error
}
