package models

type Order struct {
	ID     int     `json:"id"`
	UserId string  `json:"user_id"`
	Total  float64 `json:"total"`

	// Add the following fields
	OrderItems []OrderItem `json:"order_items"`
}
