package handlers

import (
	"net/http"

	"github.com/iyasz/JWT-RefreshToken-Go/internal/services"
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

func (ah *authHandler) Login(w http.ResponseWriter, r *http.Request){
	
}

func (ah *authHandler) Register(w http.ResponseWriter, r *http.Request){

}