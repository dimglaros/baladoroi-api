package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Game struct {
	Base
	ScheduledAt  time.Time
	TeamSize     uint
	HostID       uuid.UUID `gorm:"type:char(36)"`
	Host         User
	FieldID      uuid.UUID `gorm:"type:char(36)"`
	Participants []User    `gorm:"many2many:game_participants"`
}
