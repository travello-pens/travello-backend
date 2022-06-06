package model

type TravelAgent struct {
	ID           int    `json:"id" form:"id"`
	Name         string `json:"name" form:"name"`
	Address      string `json:"address" form:"address"`
	Phone_Number string `json:"phone_number" form:"phone_number"`
}
