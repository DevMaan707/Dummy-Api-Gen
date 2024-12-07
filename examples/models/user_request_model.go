package models

type UserRequestModel struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type ProductRequestModel struct {
	ProductName string  `json:"product_name" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Quantity    int     `json:"quantity" validate:"required,gt=0"`
}
type UserAddressRequestModel struct {
	UserID  int    `json:"user_id" validate:"required"`
	Street  string `json:"street" validate:"required"`
	City    string `json:"city" validate:"required"`
	State   string `json:"state" validate:"required"`
	ZipCode string `json:"zip_code" validate:"required"`
}
type OrderRequestModel struct {
	UserID     int     `json:"user_id" validate:"required"`
	ProductID  int     `json:"product_id" validate:"required"`
	Quantity   int     `json:"quantity" validate:"required,gt=0"`
	TotalPrice float64 `json:"total_price" validate:"required,gt=0"`
}
