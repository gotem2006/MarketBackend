package model


type Order struct {
	DelivaryAddress string `json:"delivary_address" db:"delivary_address"`
	Username string `json:"username" db:"username"`
	Phone string `json:"phone" db:"phone"`
	Email string `json:"email" db:"email"`
	Products []Product `json:"products" db:"products"`
}