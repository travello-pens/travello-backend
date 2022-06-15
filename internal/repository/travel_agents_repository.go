package repository

import (
	"fmt"

	"travel-agent-backend/internal/adapter"
	"travel-agent-backend/internal/model"

	"gorm.io/gorm"
)

func (r *RepositoryMysqlLayer) CreateTravelAgent(agent model.TravelAgent) error {
	res := r.DB.Create(&agent)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *RepositoryMysqlLayer) GetAllTravelAgents() []model.TravelAgent {
	agents := []model.TravelAgent{}
	r.DB.Find(&agents)

	return agents
}

func (r *RepositoryMysqlLayer) GetTravelAgentByID(id int) (agent model.TravelAgent, err error) {
	res := r.DB.Where("id = ?", id).Find(&agent)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *RepositoryMysqlLayer) GetSumProductByAgent(agent string) (sum int64, err error) {
	var agt model.TravelAgent
	res1 := r.DB.Where("name = ?", agent).Find(&agt)
	if res1.RowsAffected < 1 {
		err = fmt.Errorf("travel agent not found")
	}

	res2 := r.DB.Table("products").Where("id_travel_agent = ?", agt.ID).Count(&sum)
	if res2.RowsAffected < 1 {
		err = fmt.Errorf("product not found")
	}
	// res := r.DB.Where("uac.user_id = ?", "userId").Count(&count)

	return
}

func (r *RepositoryMysqlLayer) UpdateTravelAgentByID(id int, agent model.TravelAgent) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&agent)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *RepositoryMysqlLayer) DeleteTravelAgentByID(id int) error {
	res := r.DB.Delete(&model.TravelAgent{ID: id})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlTravelAgentRepository(db *gorm.DB) adapter.AdapterTravelAgentRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
