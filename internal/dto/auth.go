package dto

import (
	"github.com/iyasz/JWT-RefreshToken-Go/internal/models"
)

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthTokenResponse struct {
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	ExpiresRefreshTokenAt int    `json:"expires_refresh_token_at"`
	ExpiresAccessTokenAt  int    `json:"expires_access_token_at"`
}

type LoginResponse struct {
	AccessToken          string `json:"access_token"`
	ExpiresAccessTokenAt int    `json:"expires_access_token_at"`
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
