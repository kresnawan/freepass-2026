package models

import "github.com/oklog/ulid/v2"

type Feedback struct {
	FeedbackId  int       `json:"feedback_id"`
	OrderId     ulid.ULID `json:"order_id"`
	Description string    `json:"description"`
	Star        int       `json:"star"`
}
