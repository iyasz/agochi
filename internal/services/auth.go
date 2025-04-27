package services

import (
	"context"
	"errors"
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

	// Tipe 1 
	res, err := as.repo.FindByUsername(ctx, req.Username)
	if err != nil{
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
			return dto.RegisterResponse{}, helpers.New("", err.Error(), http.StatusRequestTimeout)
		}

		return dto.RegisterResponse{}, helpers.New("", err.Error(), http.StatusInternalServerError)
	}

	if res != nil {
		return dto.RegisterResponse{}, helpers.New("username", "username already exists", http.StatusConflict)
	}

	// Tipe 2 
	if err := as.repo.FindByEmail(ctx, req.Email); err != nil{
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
			return dto.RegisterResponse{}, helpers.New("", err.Error(), http.StatusRequestTimeout)
		}

		var httpErr *helpers.HttpError
		if errors.As(err, &httpErr) {
			if httpErr.StatusCode == http.StatusConflict {
				return dto.RegisterResponse{}, helpers.New("email", httpErr.Message, httpErr.StatusCode)
			}
		}

		return dto.RegisterResponse{}, helpers.New("", err.Error(), http.StatusInternalServerError)
	}

	user := &models.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Role:     "member",
	}

	if err := as.repo.Save(ctx, user); err != nil {
		return dto.RegisterResponse{}, helpers.New("", err.Error(), http.StatusInternalServerError)
	}

	return dto.RegisterResponse{
		User: user,
	}, nil
}
