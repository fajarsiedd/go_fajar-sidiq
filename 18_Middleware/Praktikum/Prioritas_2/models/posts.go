package models

type Posts struct {
	Base
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}
