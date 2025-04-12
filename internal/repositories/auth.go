package repositories

import (
	"context"

	"github.com/iyasz/JWT-RefreshToken-Go/internal/models"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepo(con *gorm.DB) AuthRepository {
	return &authRepository{
		db: con,
	}
}

func (ar *authRepository) Save(ctx context.Context, user *models.User) error {
	if err := ar.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (ar *authRepository) FindByEmail(ctx context.Context, user *models.User) error {
	return ar.db.Create(user).Error
}
