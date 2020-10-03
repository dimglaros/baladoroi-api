package models

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Credentials struct {
	Base
	Email    string    `json:"email"`
	Password string    `json:"password"`
	UserID   uuid.UUID `gorm:"type:char(36); unique"`
}

func (c *Credentials) BeforeSave(*gorm.DB) (err error) {
	if c.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		c.Password = string(hashedPassword)
	}
	return nil
}
