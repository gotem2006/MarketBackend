package orders

import (
	"context"
	"fmt"
	"log"
	model "shop/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type orderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *orderRepository {
	return &orderRepository{db}
}

func (o *orderRepository) AddOrder(order model.Order) error {
	validProductsQuery := `
		SELECT COUNT(*) FROM products WHERE product_name = $1 AND product_price = $2 AND category_id = $3
	`
	for _, product := range order.Products {
		var IsExist int
		o.db.QueryRow(context.Background(), validProductsQuery, product.Name, product.Price, product.Category.Id).Scan(&IsExist)
		log.Println(IsExist)
		if IsExist < 1 {
			return fmt.Errorf("No such product %v", product)
		}
	}
	query := `
		INSERT INTO orders (delivary_address, username, phone, email, products) 
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := o.db.Exec(context.Background(), query, order.DelivaryAddress, order.Username, order.Phone, order.Email, order.Products)
	return err
}

func (o *orderRepository) GetOrder(sessionKey string) (model.Order, error) {
	query := `
		SELECT delivary_address, username, phone, email, products FROM orders WHERE username = $1
	`
	var order model.Order
	err := o.db.QueryRow(context.Background(), query, sessionKey).Scan(&order.DelivaryAddress, &order.Phone, &order.Email, &order.Products)
	if err != nil{
		return order, err
	}
	return order, nil
}
