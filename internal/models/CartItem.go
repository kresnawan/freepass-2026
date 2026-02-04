package models

import "github.com/oklog/ulid/v2"

type CartItem struct {
	AccountId    ulid.ULID `json:"account_id"`
	CanteenId    int       `json:"canteen_id"`
	MenuId       int       `json:"menu_id"`
	Quantity     int       `json:"quantity"`
	PricePerItem int       `json:"price_per_item"`
}

type CartItemOrder struct {
	MenuId   int `json:"menu_id"`
	Quantity int `json:"quantity"`
}
