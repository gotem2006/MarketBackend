package basket

import "shop/internal/repository"





type basketService struct{
	repo repository.BasketRepository
}
func NewBasketService(repo repository.BasketRepository)*basketService{
	return &basketService{repo: repo}
}