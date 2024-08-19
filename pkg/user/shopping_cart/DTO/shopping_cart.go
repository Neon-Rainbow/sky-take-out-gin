package DTO

type ShoppingCartRequest struct {
	DishFlavor string `json:"dish_flavor"`
	DishID     uint   `json:"dish_id"`
	SetMealID  uint   `json:"set_meal_id"`
}
