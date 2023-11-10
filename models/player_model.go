package models

import (
	"time"
)

type Player struct {
	Id            int        `json:"id"`
	RoundId       int        `json:"roundId"`
	UserId        int        `json:"userId"`
	WalletAddress string     `json:"walletAddress"`
	BidAmount     float32    `json:"bidAmount"`
	Color         string     `json:"color"`
	CreatedAt     *time.Time `json:"createdAt,string,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,string,omitempty"`
}

func (e *Player) PlayersTable() string {
	return "players"
}
