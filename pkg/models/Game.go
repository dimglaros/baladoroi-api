package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Game struct {
	Base
	ScheduledAt  time.Time `json:"scheduledAt"`
	TeamSize     uint      `json:"teamSize"`
	HostID       uuid.UUID `gorm:"type:char(36)" json:"hostId"`
	Host         User      `json:"host"`
	FieldID      uuid.UUID `gorm:"type:char(36)" json:"fieldId"`
	Field        Field     `json:"field"`
	Participants []User    `gorm:"many2many:game_participants" json:"participants"`
}
