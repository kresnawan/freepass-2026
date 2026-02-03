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
