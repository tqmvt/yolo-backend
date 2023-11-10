package models

import (
	"time"
)

type Round struct {
	Id            int        `json:"id"`
	Round         int        `json:"round"`
	MaxDuration   int        `json:"maxDuration"`
	MaxPlayers    int        `json:"maxPlayers"`
	ActivePlayers int        `json:"activePlayers"`
	Winner        string     `json:"winner"`
	CreatedAt     *time.Time `json:"createdAt,string,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,string,omitempty"`
}

func (e *Round) RoundsTable() string {
	return "rounds"
}
