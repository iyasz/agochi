package services

import (
	"context"

	"github.com/iyasz/JWT-RefreshToken-Go/internal/config"
	"github.com/iyasz/JWT-RefreshToken-Go/internal/dto"
	"github.com/iyasz/JWT-RefreshToken-Go/internal/repositories"
)

type authService struct {
	repo   repositories.AuthRepository
	config *config.Config
}

func NewAuthService(cf *config.Config, authRepository repositories.AuthRepository) AuthService {
	return &authService{
		config: cf,
		repo:   authRepository,
	}
}

func (as *authService) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	return "", nil
}

func (as *authService) Register(ctx context.Context, req dto.RegisterRequest) error {
	return nil
}
