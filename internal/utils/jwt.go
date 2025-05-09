package utils

import (
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id string, isAdmin bool, duration time.Duration, secretKey string) (string, *TokenClaims, error) {
	newClaims, err := NewTokenClaims(id, isAdmin, duration)
	if err != nil {
		return "", nil, err
	}

	jsonData, err := json.Marshal(newClaims)
	if err != nil {
		return "", nil, err
	}
	
	var claims jwt.MapClaims
	if err := json.Unmarshal(jsonData, &claims); err != nil {
		return "", nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	tokenStr, err := token.SignedString([]byte(secretKey))
	
	if err != nil {
		return "", nil, err
	}

	return tokenStr, newClaims, nil
}