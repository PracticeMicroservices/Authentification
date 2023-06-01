package main

import (
	"authentification/cmd/api/controllers"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type App struct {
	Authentication controllers.Authentication
	DB             *sql.DB
}

func NewApp(db *sql.DB) *App {
	return &App{
		Authentication: controllers.NewAuthenticationController(db),
		DB:             db,
	}
}

func (a *App) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	mux.Use(middleware.Heartbeat("/healthCheck"))

	mux.Post("/authentication", a.Authentication.Authenticate)
	return mux
}
