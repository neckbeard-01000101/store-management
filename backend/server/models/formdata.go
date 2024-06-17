package models

type FormData struct {
	OrderNum      int    `json:"order-num"`
	OrderState    string `json:"order-state"`
	CustomerName  string `json:"customer-name"`
	CustomerCity  string `json:"customer-city"`
	CustomerPhone int    `json:"customer-phone"`
	SellerName    string `json:"seller-name"`
	TotalCost     int    `json:"total-cost"`
	SellerProfit  int    `json:"seller-profit"`
	DeliveryFee   int    `json:"delivery-fee"`
	Size          string `json:"size"`
	Color         string `json:"color"`
	ClothesType   string `json:"clothes-type"`
	ProductCost   int    `json:"cost-of-product"`
	TotalProfit   int
	SentDate      string
}
