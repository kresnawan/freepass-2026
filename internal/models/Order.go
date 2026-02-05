package models

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Order struct {
	OrderId    ulid.ULID  `json:"order_id"`
	CustomerId ulid.ULID  `json:"customer_id"`
	CanteenId  int        `json:"canteen_id"`
	Status     int        `json:"status"`
	IsPaid     bool       `json:"is_paid"`
	PaidAt     *time.Time `json:"paid_at"`
	CreatedAt  time.Time  `json:"created_at"`
}

type OrderClean struct {
	OrderId    ulid.ULID  `json:"order_id"`
	CustomerId ulid.ULID  `json:"customer_id"`
	CanteenId  int        `json:"canteen_id"`
	Status     string     `json:"status"`
	IsPaid     bool       `json:"is_paid"`
	PaidAt     *time.Time `json:"paid_at"`
	CreatedAt  time.Time  `json:"created_at"`
}

type OrderDetails struct {
	OrderClean
	TotalPrice int         `json:"total_price"`
	Items      []OrderItem `json:"items"`
}

type OrderItem struct {
	MenuId       int `json:"menu_id"`
	Quantity     int `json:"quantity"`
	PricePerItem int `json:"price_per_item"`
}
