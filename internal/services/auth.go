package services

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/iyasz/JWT-RefreshToken-Go/internal/config"
	"github.com/iyasz/JWT-RefreshToken-Go/internal/dto"
	"github.com/iyasz/JWT-RefreshToken-Go/internal/helpers"
	"github.com/iyasz/JWT-RefreshToken-Go/internal/models"
	"github.com/iyasz/JWT-RefreshToken-Go/internal/repositories"
	"github.com/iyasz/JWT-RefreshToken-Go/internal/utils"
	"golang.org/x/crypto/bcrypt"
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

func (as *authService) Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthTokenResponse, error) {

	res, err := as.repo.FindByUsername(ctx, req.Username)

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, helpers.New("", err.Error(), http.StatusRequestTimeout)
		}
		if errors.Is(err, context.Canceled) {
			return nil, helpers.New("", err.Error(), http.StatusBadRequest)
		}

		return nil, helpers.New("", err.Error(), http.StatusInternalServerError)
	}

	if res == nil {
		return nil, helpers.New("", "Invalid username or password", http.StatusUnauthorized)
	}

	// Verify Password
	errCompare := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(req.Password))
	if errCompare != nil {
		return nil, helpers.New("", "Invalid username or password", http.StatusUnauthorized)
	}

	// Generate Access Token
	accessToken, accessClaims, err := utils.GenerateToken(res.ID.String(), res.Role == "admin", 15*time.Minute, as.config.Jwt.AccessKey)
	if err != nil {
		utils.Log.Error("Failed to generate access token", "Error", err)
		return nil, helpers.New("", "Error while generating access token", http.StatusUnauthorized)
	}

	// Generate Refresh Token
	refreshToken, refreshClaims, err := utils.GenerateToken(res.ID.String(), res.Role == "admin", 7*24*time.Hour, as.config.Jwt.AccessKey)
	if err != nil {
		utils.Log.Error("Failed to generate refresh token", "Error", err)
		return nil, helpers.New("", "Error while generating refresh token", http.StatusUnauthorized)
	}
	
	return &dto.AuthTokenResponse{
		AccessToken:          accessToken,
		RefreshToken:         refreshToken,
		ExpiresRefreshTokenAt: refreshClaims.Exp,
		ExpiresAccessTokenAt: accessClaims.Exp,
	}, nil

}

func (as *authService) Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error) {

	// Tipe 1
	res, err := as.repo.FindByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return dto.RegisterResponse{}, helpers.New("", err.Error(), http.StatusRequestTimeout)
		}
		if errors.Is(err, context.Canceled) {
			return dto.RegisterResponse{}, helpers.New("", err.Error(), http.StatusBadRequest)
		}

		return dto.RegisterResponse{}, helpers.New("", err.Error(), http.StatusInternalServerError)
	}

	if res != nil {
		return dto.RegisterResponse{}, helpers.New("username", "username already exists", http.StatusConflict)
	}

	// Tipe 2
	if err := as.repo.FindByEmail(ctx, req.Email); err != nil {
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

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash password ", err)
		return dto.RegisterResponse{}, helpers.New("", err.Error(), http.StatusInternalServerError)
	}

	user := &models.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "member",
	}

	if err := as.repo.Save(ctx, user); err != nil {
		return dto.RegisterResponse{}, helpers.New("", err.Error(), http.StatusInternalServerError)
	}

	return dto.RegisterResponse{
		User: user,
	}, nil
}
