package request

import "go-wishlist-api/models"

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l LoginInput) ToModel() models.Users {
	return models.Users{
		Email:    l.Email,
		Password: l.Password,
	}
}
