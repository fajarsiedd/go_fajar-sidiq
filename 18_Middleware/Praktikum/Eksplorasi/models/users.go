package models

type Users struct {
	Base
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Role     string `json:"role" gorm:"default:'user'"`
	Password string `json:"password"`
}
