package basket

import "github.com/jackc/pgx/v5/pgxpool"




type basketRepository struct{
	db *pgxpool.Pool
}



func NewBasketRepository(db *pgxpool.Pool) *basketRepository{
	return &basketRepository{db}
}


