package models

import "github.com/google/uuid"

type Item struct {
	ItemID      int      `json:"item_id"`
	Name        string   `json:"name"`
	Vendor      string   `json:"vendor"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
	ThreeLink   string   `json:"three_link"`
	Price       float32  `json:"price"`
	Quantity    int      `json:"quantity"`
}

type CartItem struct {
	ItemID   int `json:"item_id"`
	Quantity int `json:"quantity"`
}

type User struct {
	UserID    uuid.UUID  `json:"user_id"`
	Cart      []CartItem `json:"cart"`
	CartValue float32    `json:"total_cart_value"`
}

// It should need a timestamp for when
// Right now, the User and Order structs look the same, but on addition of timestamps, orderId, etc. fields, they would be different
// Also, Order.OrderValue would be different from User.CartValue due to possible discounts
type Order struct {
	OrderID    int        `json:"order_id"`
	UserID     uuid.UUID  `json:"user_id"`
	Items      []CartItem `json:"items"`
	OrderValue float32    `json:"order_value"`
	Discount   float32    `json:"discount"`
}

type Coupon struct {
	UserID     uuid.UUID `json:"user_id"`
	CouponCode uuid.UUID `json:"coupon_code"`
}
