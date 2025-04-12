package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/iyasz/JWT-RefreshToken-Go/internal/handlers"
)

func AuthRoutes(r *chi.Mux, authHandler handlers.AuthHandler) {

	r.Route("/auth", func(r chi.Router) {
		r.Get("/greeting", authHandler.Greeting)
		r.Post("/register", authHandler.Register)
		r.Get("/login", authHandler.Login)
	})

}