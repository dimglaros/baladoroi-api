package models

import uuid "github.com/satori/go.uuid"

const (
	Pending = iota
	Confirmed
	Canceled
)

type Participation struct {
	Base
	UserID uuid.UUID `gorm:"type:char(36)" json:"userId"`
	GameID uuid.UUID `gorm:"type:char(36)" json:"gameId"`
	Status uint8     `json:"status"`
}
