package model

import "time"

type Product struct {
	ID                 int       `json:"id" form:"id"`
	ID_Travel_Agent    int       `json:"id_travel_agent" form:"id_travel_agent"`
	ID_Location        int       `json:"id_location" form:"id_location"`
	Name_Product       string    `json:"name_product" form:"name_product"`
	Price              float64   `json:"price" form:"price"`
	Description        string    `json:"description" form:"description"`
	Length_Of_Stay     int       `json:"length_of_stay" form:"length_of_stay"`
	Address_Of_Product string    `json:"address_of_product" form:"address_of_product"`
	Tax                int       `json:"tax" form:"tax"`
	Lodging            int       `json:"lodging" form:"lodging"`
	Airport_Pickup     int       `json:"airport_pickup" form:"airport_pickup"`
	Refund             int       `json:"refund" form:"refund"`
	Created_At         time.Time `json:"created_at" form:"created_at"`
	Updated_At         time.Time `json:"updated_at" form:"updated_at"`
}
