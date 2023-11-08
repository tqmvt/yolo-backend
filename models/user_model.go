package models

import (
	"time"
)

type User struct {
	Id            int        `json:"id"`
	FirstName     string     `json:"firstName"`
	LastName      string     `json:"lastName"`
	WalletAddress string     `json:"walletAddress" binding:"required"`
	Email         string     `json:"email"`
	CreatedAt     *time.Time `json:"createdAt,string,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,string,omitempty"`
}

func (e *User) UsersTable() string {
	return "users"
}
