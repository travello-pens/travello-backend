package model

import "time"

type TravelAgent struct {
	ID           int       `json:"id" form:"id"`
	Name_Agent   string    `json:"name_agent" form:"name_agent"`
	Address      string    `json:"address" form:"address"`
	Phone_Number string    `json:"phone_number" form:"phone_number"`
	Created_At   time.Time `json:"created_at" form:"created_at"`
	Updated_At   time.Time `json:"updated_at" form:"updated_at"`
}

type TravelAgentProduct struct {
	Name_Agent         string  `json:"name_agent" form:"name_agent"`
	Address            string  `json:"address" form:"address"`
	Phone_Number       string  `json:"phone_number" form:"phone_number"`
	Name_Product       string  `json:"name_product" form:"name_product"`
	Name_Location      string  `json:"name_location" form:"name_location"`
	Price              float64 `json:"price" form:"price"`
	Description        string  `json:"description" form:"description"`
	Length_Of_Stay     int     `json:"length_of_stay" form:"length_of_stay"`
	Address_Of_Product string  `json:"address_of_product" form:"address_of_product"`
	Tax                int     `json:"tax" form:"tax"`
	Lodging            int     `json:"lodging" form:"lodging"`
	Airport_Pickup     int     `json:"airport_pickup" form:"airport_pickup"`
	Refund             int     `json:"refund" form:"refund"`
}
