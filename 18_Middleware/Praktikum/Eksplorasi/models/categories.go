package models

type Categories struct {
	Base
	Name  string  `json:"name" validate:"required"`
	Posts []Posts `json:"posts,omitempty" gorm:"foreignKey:CategoryID"`
}
