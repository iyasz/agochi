package dto

import "github.com/iyasz/JWT-RefreshToken-Go/internal/models"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email,max=250"`
	Password string `json:"password" validate:"required,max=250"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,max=250"`
	Username string `json:"username" validate:"required,max=250"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=250"`
	Role     string `json:"role"`
}

type RegisterResponse struct {
	User *models.User `json:"user"`
}

