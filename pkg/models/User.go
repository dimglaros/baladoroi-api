package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Base
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Games    []Game `gorm:"many2many:game_participants" json:"games"`
}

func (u *User) BeforeSave(*gorm.DB) (err error) {
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		u.Password = string(hashedPassword)
	}
	return nil
}
