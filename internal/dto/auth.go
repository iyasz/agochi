package dto

type LoginRequest struct{
	Email    string `json:"email" validate:"required,email,max=250"`
	Password string `json:"password" validate:"required,max=250"`
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
