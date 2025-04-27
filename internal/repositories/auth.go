package repositories

import (
	"context"
	"errors"
	"net/http"

	"github.com/iyasz/JWT-RefreshToken-Go/internal/helpers"
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
	return ar.db.WithContext(ctx).Create(user).Error
}

func (ar *authRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {

	var user models.User
	if err := ar.db.WithContext(ctx).Where("username= ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		
		return nil, err
	}

	return &user, nil
}

func (ar *authRepository) FindByEmail(ctx context.Context, email string) error {

	var user models.User
	if err := ar.db.WithContext(ctx).Where("email= ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		return err
	}

	return helpers.New("", "email already exists", http.StatusConflict)
}
