package models

import "github.com/google/uuid"

type Item struct {
	ItemID      int      `json:"id"`
	Name        string   `json:"name"`
	Vendor      string   `json:"vendor"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
	ThreeLink   string   `json:"three_link"`
	Price       float32  `json:"price"`
	Quantity    int      `json:"quantity"`
}

type CartItem struct {
	ItemID   int `json:"id"`
	Quantity int `json:"quantity"`
}

type User struct {
	UserID    uuid.UUID  `json:"id"`
	Cart      []CartItem `json:"cart"`
	CartValue float32    `json:"total_cart_value"`
}

// It should need a timestamp for when
type Order struct {
	UserID     uuid.UUID `json:"user_id"`
	Items      []Item    `json:"items"`
	OrderValue float32   `json:"order_value"`
}
