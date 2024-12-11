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

type User struct {
	UserID uuid.UUID `json:"id"`
	Cart   []Item    `json:"cart"`
}

// TODO: Have to create the Order type
type Order struct{}
