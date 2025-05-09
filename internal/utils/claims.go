package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenClaims struct {
	UserID   string `json:"sub"`
	JwtID    string `json:"jti"`
	Exp      int  `json:"exp"`
	Iat      int  `json:"iat"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

func NewTokenClaims(id string, isAdmin bool, duration time.Duration) (*TokenClaims, error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}	

	return &TokenClaims{
		UserID: id,
		JwtID:  newUUID.String(),
		Exp:    int(time.Now().Add(duration).UnixMilli()),
		Iat:    int(time.Now().UnixMilli()),
		IsAdmin:  isAdmin,
	}, nil

}
