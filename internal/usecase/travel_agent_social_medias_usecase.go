package usecase

import (
	"travel-agent-backend/config"
	"travel-agent-backend/internal/adapter"
	"travel-agent-backend/internal/model"
)

type serviceTravelAgentSocialMedia struct {
	c    config.Config
	repo adapter.AdapterTravelAgentSocialMediaRepository
}

func (s *serviceTravelAgentSocialMedia) CreateTravelAgentSocialMediaService(media model.TravelAgentSocialMedia) error {
	return s.repo.CreateTravelAgentSocialMedia(media)
}

func (s *serviceTravelAgentSocialMedia) GetAllTravelAgentSocialMediasService() []model.TravelAgentSocialMedia {
	return s.repo.GetAllTravelAgentSocialMedias()
}

func (s *serviceTravelAgentSocialMedia) GetTravelAgentSocialMediaByIDService(id int) (media model.TravelAgentSocialMedia, err error) {
	return s.repo.GetTravelAgentSocialMediaByID(id)
}

func (s *serviceTravelAgentSocialMedia) UpdateTravelAgentSocialMediaByIDService(id int, media model.TravelAgentSocialMedia) error {
	return s.repo.UpdateTravelAgentSocialMediaByID(id, media)
}

func (s *serviceTravelAgentSocialMedia) DeleteTravelAgentSocialMediaByIDService(id int) error {
	return s.repo.DeleteTravelAgentSocialMediaByID(id)
}

func NewServiceTravelAgentSocialMedia(repo adapter.AdapterTravelAgentSocialMediaRepository, c config.Config) adapter.AdapterTravelAgentSocialMediaService {
	return &serviceTravelAgentSocialMedia{
		repo: repo,
		c:    c,
	}
}
