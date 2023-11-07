package models

import (
	"time"
)

type User struct {
	Id            int        `json:"id"`
	FirstName     string     `json:"first_name"`
	LastName      string     `json:"last_name"`
	WalletAddress string     `json:"wallet_address" binding:"required"`
	Email         string     `json:"email"`
	CreatedAt     *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at_at,string,omitempty"`
}

func (e *User) UsersTable() string {
	return "users"
}
