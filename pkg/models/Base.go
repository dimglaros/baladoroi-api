package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        uuid.UUID  `gorm:"type:char(36);primary_key;" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}

func (b *Base) BeforeCreate(*gorm.DB) (err error) {
	b.ID = uuid.NewV4()
	return nil
}
