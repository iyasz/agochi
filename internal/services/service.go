package services

import (
	"context"

	"github.com/iyasz/JWT-RefreshToken-Go/internal/dto"
)


type AuthService interface{
	Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthTokenResponse, error)
	Register(ctx context.Context, req dto.RegisterRequest ) (dto.RegisterResponse, error)
}
