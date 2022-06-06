package model

type ProductPhoto struct {
	ID         int    `json:"id" form:"id"`
	ID_Product int    `json:"id_product" form:"id_product"`
	Photo_Name string `json:"photo_name" form:"photo_name"`
	File_Name  string `json:"file_name" form:"file_name"`
}
