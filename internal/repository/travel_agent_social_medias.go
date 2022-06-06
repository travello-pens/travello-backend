package repository

import (
	"fmt"

	"travel-agent-backend/internal/adapter"
	"travel-agent-backend/internal/model"

	"gorm.io/gorm"
)

func (r *RepositoryMysqlLayer) CreateTravelAgentSocialMedia(media model.TravelAgentSocialMedia) error {
	res := r.DB.Create(&media)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *RepositoryMysqlLayer) GetAllTravelAgentSocialMedias() []model.TravelAgentSocialMedia {
	medias := []model.TravelAgentSocialMedia{}
	r.DB.Find(&medias)

	return medias
}

func (r *RepositoryMysqlLayer) GetTravelAgentSocialMediaByID(id int) (media model.TravelAgentSocialMedia, err error) {
	res := r.DB.Where("id = ?", id).Find(&media)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *RepositoryMysqlLayer) UpdateTravelAgentSocialMediaByID(id int, media model.TravelAgentSocialMedia) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&media)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *RepositoryMysqlLayer) DeleteTravelAgentSocialMediaByID(id int) error {
	res := r.DB.Delete(&model.TravelAgentSocialMedia{ID: id})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlTravelAgentSocialMediaRepository(db *gorm.DB) adapter.AdapterTravelAgentSocialMediaRepository {
	return &RepositoryMysqlLayer{
		DB: db,
	}
}
