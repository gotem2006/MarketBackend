package model



type Basket struct{
	ProductId int `json:"product_id" db:"product_id"`
	Quantity int `json:"quantity" db:"quantity"`
	SessionKey string `json:"session_key" db:"session_key"`
}
