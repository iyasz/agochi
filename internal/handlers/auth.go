package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/iyasz/JWT-RefreshToken-Go/internal/dto"
	"github.com/iyasz/JWT-RefreshToken-Go/internal/helpers"
	"github.com/iyasz/JWT-RefreshToken-Go/internal/services"
	"github.com/iyasz/JWT-RefreshToken-Go/internal/utils"
)

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) AuthHandler {
	return &authHandler{
		authService: authService,
	}
}

func (ah *authHandler) Greeting(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello World"}`))
}

func (ah *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	res, err := ah.authService.Login(ctx, req)

	if err != nil {
		utils.Log.Error("Failed to login", "Error", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

func (ah *authHandler) Register(w http.ResponseWriter, r *http.Request){
	ctx, cancel := context.WithTimeout(r.Context(), 7*time.Second)
	defer cancel()

	var req dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	res, err := ah.authService.Register(ctx, req);

	if err != nil {
		if httpErr, ok := err.(*helpers.HttpError); ok {
			http.Error(w, httpErr.Message, httpErr.StatusCode)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}