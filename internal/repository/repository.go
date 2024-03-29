package repository

import (
	model "shop/internal/models"
	"shop/internal/repository/basket"
	"shop/internal/repository/product"

	"github.com/jackc/pgx/v5/pgxpool"
)




type ProductRepository interface{
	GetProduct(Id int) model.Product
	GetProducts() []model.Product
	AddProduct(product model.Product) error
	AddAttribute(attribute model.Attribute, product_id int) error
	AddCategory(category model.Category) error
}


type BasketRepository interface{

}


type Repositories struct{
	Product ProductRepository
	Basket BasketRepository
}


func NewRepositories(db *pgxpool.Pool) *Repositories{
	return &Repositories{
		product.NewProductRepository(db),
		basket.NewBasketRepository(db),
	}
}