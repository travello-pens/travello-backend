package usecase

import (
	"travel-agent-backend/config"
	"travel-agent-backend/internal/adapter"
	"travel-agent-backend/internal/model"
)

type serviceTravelAgent struct {
	c    config.Config
	repo adapter.AdapterTravelAgentRepository
}

func (s *serviceTravelAgent) CreateTravelAgentService(agent model.TravelAgent) error {
	return s.repo.CreateTravelAgent(agent)
}

func (s *serviceTravelAgent) GetAllTravelAgentsService() []model.TravelAgent {
	return s.repo.GetAllTravelAgents()
}

func (s *serviceTravelAgent) GetTravelAgentByIDService(id int) (agent model.TravelAgent, err error) {
	return s.repo.GetTravelAgentByID(id)
}

func (s *serviceTravelAgent) GetSumProductByAgentService(agent string) (sum int64, err error) {
	return s.repo.GetSumProductByAgent(agent)
}

func (s *serviceTravelAgent) GetProductsByAgentService() ([]model.TravelAgentProduct, error) {
	return s.repo.GetProductsByAgentTravel()
}

func (s *serviceTravelAgent) UpdateTravelAgentByIDService(id int, agent model.TravelAgent) error {
	return s.repo.UpdateTravelAgentByID(id, agent)
}

func (s *serviceTravelAgent) DeleteTravelAgentByIDService(id int) error {
	return s.repo.DeleteTravelAgentByID(id)
}

func NewServiceTravelAgent(repo adapter.AdapterTravelAgentRepository, c config.Config) adapter.AdapterTravelAgentService {
	return &serviceTravelAgent{
		repo: repo,
		c:    c,
	}
}
