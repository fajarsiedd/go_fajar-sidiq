package response

import "rest/models"

type LoginResponse struct {
	models.Base
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}
