package handlers

import (
	"context"
	"encoding/json"
	"log"
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
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()


	var req dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if validate := utils.Validate(req); len(validate) > 0 {
		helpers.NewErrorResponse(w, http.StatusUnprocessableEntity, "Validation Error", "", validate)
		return
	}

	if ok := utils.NoSpace(req.Username); !ok {
		helpers.NewErrorResponse(w, http.StatusUnprocessableEntity, "Validation Error", "", map[string]string{"username": "username cannot contain spaces"})
		return
	}

	if ok := utils.NoSpace(req.Password); !ok {
		helpers.NewErrorResponse(w, http.StatusUnprocessableEntity, "Validation Error", "", map[string]string{"password": "password cannot contain spaces"})
		return	
	}

	res, err := ah.authService.Register(ctx, req);

	if err != nil {
		if httpErr, ok := err.(*helpers.HttpError); ok {
			if httpErr.Field != "" {
				helpers.NewErrorResponse(w, httpErr.StatusCode, "Validation Error", "", map[string]string{httpErr.Field: httpErr.Message})
				return
			}

			helpers.NewErrorResponse(w, httpErr.StatusCode, httpErr.Message, "", nil)
			return
		}

		log.Println(err.Error())
		helpers.NewErrorResponse(w, http.StatusInternalServerError, "Internal Server Error", "", nil)
		return
	}

	helpers.NewSuccessResponse(w, http.StatusOK, "Resource has been successfully created.", res, nil)
}		