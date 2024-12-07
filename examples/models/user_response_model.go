package models

type UserResponseModel struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}
type ProductResponseModel struct {
	ID          int     `json:"id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	CreatedAt   string  `json:"created_at"`
}
type UserAddressResponseModel struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zip_code"`
	CreatedAt string `json:"created_at"`
}
type OrderResponseModel struct {
	ID         int     `json:"id"`
	UserID     int     `json:"user_id"`
	ProductID  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
	CreatedAt  string  `json:"created_at"`
}
