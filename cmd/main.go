package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/iyasz/JWT-RefreshToken-Go/internal/api"
	"github.com/iyasz/JWT-RefreshToken-Go/internal/config"
	"github.com/iyasz/JWT-RefreshToken-Go/internal/connection"
	"github.com/iyasz/JWT-RefreshToken-Go/internal/handlers"
	"github.com/iyasz/JWT-RefreshToken-Go/internal/repositories"
	"github.com/iyasz/JWT-RefreshToken-Go/internal/services"
)

func main() {

	cf := config.LoadDB()
	con := connection.GetConnection(cf.Database)

	// Setup Repository 
	authRepository := repositories.NewAuthRepo(con)

	// Setup Service 
	authService := services.NewAuthService(cf, authRepository)

	// Setup Handler 
	authHandler := handlers.NewAuthHandler(authService)

	r := chi.NewRouter()

	// Setup Routes 
	r.Route("/api", func(r chi.Router) {
		api.AuthRoutes(r, authHandler)
	})

	
	log.Println("server running on http://" + cf.Server.Host + ":" + cf.Server.Port)
	http.ListenAndServe(":"+cf.Server.Port, r)
}