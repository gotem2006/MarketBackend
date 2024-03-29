package basket

import "shop/internal/service"





type basketHandler struct{
	BasketServcie service.BasketService
}



func NewBasketHandler(basketService service.BasketService) *basketHandler{
	return &basketHandler{
		BasketServcie: basketService,
	}
}