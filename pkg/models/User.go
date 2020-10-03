package models

type User struct {
	Base
	Name        string      `json:"name"`
	Surname     string      `json:"surname"`
	Phone       string      `json:"phone"`
	Credentials Credentials `gorm:"constraint:OnUpdate:CASCADE" ;json:"credentials"`
	Games       []Game      `gorm:"many2many:game_participants"`
}
