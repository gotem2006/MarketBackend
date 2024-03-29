package model




type Attribute struct{
	Name string `json:"name" db:"attribute_name"`
	Value string `json:"value" db:"attribute_value"`
}