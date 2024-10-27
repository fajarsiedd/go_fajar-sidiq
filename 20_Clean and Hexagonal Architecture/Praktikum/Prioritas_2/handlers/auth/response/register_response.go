package response

import "go-wishlist-api/models"

type RegisterResponse struct {
	models.Base
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Register_FromModel(u models.Users) RegisterResponse {
	return RegisterResponse{
		Base: models.Base{
			ID:        u.ID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
			DeletedAt: u.DeletedAt,
		},
		Name:  u.Name,
		Email: u.Email,
	}
}
