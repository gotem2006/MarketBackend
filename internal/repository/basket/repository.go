package basket

import (
	"context"
	model "shop/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type basketRepository struct {
	db *pgxpool.Pool
}

func NewBasketRepository(db *pgxpool.Pool) *basketRepository {
	return &basketRepository{db}
}

func (b *basketRepository) AddBasket(basket model.Basket) error {
	query := `
		INSERT INTO baskets (quantity, session_key, product_id) VALUES ($1, $2, $3)
		ON CONFLICT (session_key, product_id) DO UPDATE 
		SET quantity = baskets.quantity + $1
	`
	_, err := b.db.Exec(context.Background(), query, basket.Quantity, basket.SessionKey, basket.ProductId)
	return err
}

func (b *basketRepository) GetBasket(sessionKey string) (Baskets []model.Basket, err error) {
	query := `
		SELECT product_id, quantity, session_key FROM baskets WHERE session_key = $1
	`
	rows, err := b.db.Query(context.Background(), query, sessionKey)
	if err != nil {
		return nil, err
	}
	Baskets, err = pgx.CollectRows(rows, pgx.RowToStructByName[model.Basket])
	if err != nil {
		return nil, err
	}
	return Baskets, nil
}

func (b *basketRepository) ReduceQuantity(basket model.Basket) error {
	query := `
		UPDATE baskets SET quantity = baskets.quantity - $1 WHERE session_key = $2 AND product_id = $3 RETUTNING quantity
	`
	var quantity int
	b.db.QueryRow(context.Background(), query, basket.Quantity, basket.SessionKey, basket.ProductId).Scan(&quantity)
	if quantity < 1 {
		query = "DELETE FROM baskets WHERE session_key = $1 AND product_id = $2"
		_, err := b.db.Exec(context.Background(), query, basket.SessionKey, basket.ProductId)
		if err != nil{
			return err
		}
	}
	return nil
}

func (b *basketRepository) DeleteBasket(basket model.Basket) error {
	query := "DELETE FROM baskets WHERE session_key = $1 AND product_id = $2"
	_, err := b.db.Exec(context.Background(), query, basket.SessionKey, basket.ProductId)
	return err
}
