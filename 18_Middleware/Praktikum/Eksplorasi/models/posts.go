package models

type Posts struct {
	Base
	Title      string `json:"title" validate:"required"`
	Content    string `json:"content" validate:"required"`
	CategoryID string `json:"category_id" gorm:"size:191"`
}
