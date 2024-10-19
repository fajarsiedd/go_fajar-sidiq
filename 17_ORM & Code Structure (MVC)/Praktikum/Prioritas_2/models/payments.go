package models

type Payments struct {
	ID          string `json:"id" form:"id" gorm:"primaryKey"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Amount      int    `json:"amount" form:"amount"`
}
