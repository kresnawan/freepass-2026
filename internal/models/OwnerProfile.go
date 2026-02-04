package models

import "github.com/oklog/ulid/v2"

type OwnerProfile struct {
	AccountId ulid.ULID `json:"account_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Role      string    `json:"role"`
}
