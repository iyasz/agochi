package repositories

import (
	"context"

	"github.com/iyasz/JWT-RefreshToken-Go/internal/models"
)

type AuthRepository interface {
	Save(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) error
	FindByUsername(ctx context.Context, username string) error
}