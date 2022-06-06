package model

type TravelAgentSocialMedia struct {
	ID              int    `json:"id" form:"id"`
	ID_Travel_Agent int    `json:"id_travel_agent" form:"id_travel_agent"`
	Facebook        string `json:"facebook" form:"facebook"`
	Twitter         string `json:"twitter" form:"twitter"`
	Instagram       string `json:"instagram" form:"instagram"`
	Youtube         string `json:"youtube" form:"youtube"`
	Tiktok          string `json:"tiktok" form:"tiktok"`
	Website         string `json:"website" form:"website"`
}
