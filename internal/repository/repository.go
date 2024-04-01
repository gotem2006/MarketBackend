package repository

import (
	model "shop/internal/models"
	"shop/internal/repository/basket"
	"shop/internal/repository/orders"
	"shop/internal/repository/product"

	"github.com/jackc/pgx/v5/pgxpool"
)




type ProductRepository interface{
	GetProduct(Id int) (product model.Product, err error)
	GetProducts() (Products []model.Product, err error)
	AddProduct(product model.Product) error
	AddAttribute(attribute model.Attribute, product_id int) error
	AddCategory(category model.Category) error
}


type BasketRepository interface{
	AddBasket(basket model.Basket) error
	GetBasket(sessionKey string) (Baskets []model.Basket, err error)
	ReduceQuantity(basket model.Basket) error
	DeleteBasket(basket model.Basket) error
}


type OrderRepository interface{
	AddOrder(order model.Order) error
	GetOrder(sessionKey string) (model.Order, error)
}


type Repositories struct{
	Product ProductRepository
	Basket BasketRepository
	Order OrderRepository
}


func NewRepositories(db *pgxpool.Pool) *Repositories{
	return &Repositories{
		product.NewProductRepository(db),
		basket.NewBasketRepository(db),
		orders.NewOrderRepository(db),
	}
}