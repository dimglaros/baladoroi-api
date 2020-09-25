package models

type User struct {
	Base
	Name        string
	Surname     string
	Phone       string
	Credentials Credentials
	Games       []Game `gorm:"many2many:game_participants"`
}
