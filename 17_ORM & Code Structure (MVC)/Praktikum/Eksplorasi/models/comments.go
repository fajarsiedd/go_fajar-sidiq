package models

type Comments struct {
	Base
	Content string `json:"content"`
	PostID  string `json:"post_id" gorm:"size:191"`
}
