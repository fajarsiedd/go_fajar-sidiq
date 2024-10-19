package models

type Posts struct {
	Base
	Title    string     `json:"title"`
	Content  string     `json:"content"`
	Comments []Comments `json:"comments,omitempty" gorm:"foreignKey:PostID"`
}
