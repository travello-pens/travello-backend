package model

type Location struct {
	ID         int    `json:"id" form:"id"`
	ID_Product int    `json:"id_product" form:"id_product"`
	Name       string `json:"name" form:"name"`
}
