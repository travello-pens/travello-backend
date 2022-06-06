package adapter

import "travel-agent-backend/internal/model"

type AdapterTravelAgentRepository interface {
	CreateTravelAgent(agent model.TravelAgent) error
	GetAllTravelAgents() []model.TravelAgent
	GetTravelAgentByID(id int) (agent model.TravelAgent, err error)
	UpdateTravelAgentByID(id int, agent model.TravelAgent) error
	DeleteTravelAgentByID(id int) error
}

type AdapterTravelAgentService interface {
	CreateTravelAgentService(agent model.TravelAgent) error
	GetAllTravelAgentsService() []model.TravelAgent
	GetTravelAgentByIDService(id int) (agent model.TravelAgent, err error)
	UpdateTravelAgentByIDService(id int, agent model.TravelAgent) error
	DeleteTravelAgentByIDService(id int) error
}
