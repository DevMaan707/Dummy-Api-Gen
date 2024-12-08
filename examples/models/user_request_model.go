package models

type UserRequestModel struct {
	UserID string `json:"user_id"`
}

type ProductRequestModel struct {
	ProductID string `json:"product_id"`
}
type UserAddressRequestModel struct {
	UserID int `json:"user_id" validate:"required"`
}
type OrderRequestModel struct {
	UserID     int     `json:"user_id" validate:"required"`
	ProductID  int     `json:"product_id" validate:"required"`
	Quantity   int     `json:"quantity" validate:"required,gt=0"`
	TotalPrice float64 `json:"total_price" validate:"required,gt=0"`
}
type OrderDetailsRequestModel struct {
	OrderID string `json:"order_id"`
}
