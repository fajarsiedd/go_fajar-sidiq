package models

type Packages struct {
	ID               string  `json:"id" form:"id" gorm:"primaryKey"`
	Name             string  `json:"name" form:"name"`
	SenderName       string  `json:"sender_name" form:"sender_name"`
	ReceiverName     string  `json:"receiver_name" form:"receiver_name"`
	SenderLocation   string  `json:"sender_location" form:"sender_location"`
	ReceiverLocation string  `json:"receiver_location" form:"receiver_location"`
	Fee              int     `json:"fee" form:"fee"`
	Weight           float64 `json:"weight" form:"weight"`
}
