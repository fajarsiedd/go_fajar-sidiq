package models

type Packages struct {
	Base
	Name             string  `json:"name" validate:"required"`
	Sender           string  `json:"sender" validate:"required"`
	Receiver         string  `json:"receiver" validate:"required"`
	SenderLocation   string  `json:"sender_location" validate:"required"`
	ReceiverLocation string  `json:"receiver_location" validate:"required"`
	Fee              int     `json:"fee" validate:"required"`
	Weight           float64 `json:"weight" validate:"required"`
}
