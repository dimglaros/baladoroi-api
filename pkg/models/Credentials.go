package models

import uuid "github.com/satori/go.uuid"

type Credentials struct {
	Base
	Email    string
	Password string
	UserID   uuid.UUID `gorm:"type:char(36)"`
}
