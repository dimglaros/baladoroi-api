package models

type Field struct {
	Base
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Games   []Game `json:"games"`
}
