package request

import "go-wishlist-api/models"

type RegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r RegisterInput) ToModel() models.Users {
	return models.Users{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
	}
}
