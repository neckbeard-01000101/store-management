package models

type Storage struct {
	TypeOfProduct string `json:"type_of_product"`
	Quantity      int    `json:"quantity"`
	Color         string `json:"product_color"`
	Size          string `json:"product_size"`
}
