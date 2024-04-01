package model



type Product struct{
	Name string `json:"name" db:"product_name"`
	Price int `json:"price" db:"product_price"`
	Category Category
	Attributes []Attribute
}