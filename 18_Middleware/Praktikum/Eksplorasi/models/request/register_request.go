package request

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Role     string `json:"role"`
	Password string `json:"password" validate:"required"`
}
