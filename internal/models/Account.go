package models

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type Account struct {
	AccountId   ulid.ULID  `json:"account_id"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	Password    string     `json:"password"`
	Role        string     `json:"role"`
	IsActive    bool       `json:"is_active"`
	DeactivedAt *time.Time `json:"deactived_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
