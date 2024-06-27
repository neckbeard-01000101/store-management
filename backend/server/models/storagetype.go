package models

type Storage struct {
	TypeOfProduct string `bson:"type_of_product"`
	Quantity      int    `bson:"quantity"`
}
