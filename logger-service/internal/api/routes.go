package routes

import (
	"net/http"

	"github.com/Dimoonevs/SportsApp/logger-service/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type AppLogger struct {
	Service *service.LogsService
}

func (app *AppLogger) Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/ping"))

	mux.Post("/logs", app.WriteLogs)

	return mux
}

func (app *AppLogger) WriteLogs(w http.ResponseWriter, r *http.Request) {
	app.Service.WriteLogs(r.Context(), r)
}
