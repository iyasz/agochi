package services

import (
	"context"
	"net/http"

	"github.com/iyasz/JWT-RefreshToken-Go/internal/config"
	"github.com/iyasz/JWT-RefreshToken-Go/internal/dto"
	"github.com/iyasz/JWT-RefreshToken-Go/internal/helpers"
	"github.com/iyasz/JWT-RefreshToken-Go/internal/models"
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

func (as *authService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	return dto.LoginResponse{}, nil
}

func (as *authService) Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error) {
	if err := as.repo.FindByEmail(ctx, req.Email); err == nil{
		return dto.RegisterResponse{}, helpers.New("email already exists", http.StatusConflict)
	}

	if err := as.repo.FindByUsername(ctx, req.Username); err == nil{
		return dto.RegisterResponse{}, helpers.New("username already exists", http.StatusConflict)
	}

	user := &models.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Role:     "member",
	}

	if err := as.repo.Save(ctx, user); err != nil {
		return dto.RegisterResponse{}, helpers.New(err.Error(), http.StatusInternalServerError)
	}

	return dto.RegisterResponse{
		User: user,
	}, nil
}
