package response

import "go-wishlist-api/models"

type LoginResponse struct {
	models.Base
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

func Login_FromModel(u models.Users, token string) LoginResponse {
	return LoginResponse{
		Base: models.Base{
			ID:        u.ID,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
			DeletedAt: u.DeletedAt,
		},
		Name:        u.Name,
		Email:       u.Email,
		AccessToken: token,
	}
}
