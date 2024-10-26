package response

import "rest/models"

type RegisterResponse struct {
	models.Base
	Name  string `json:"name"`
	Email string `json:"email"`
}
