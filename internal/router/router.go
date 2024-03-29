package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/tjingsheng/jsask-forum-backend/internal/routes"
)

func Setup() chi.Router {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedOrigins:   []string{"https://www.jsask.online", "https://www.jsask.site"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	setUpRoutes(r)
	return r
}

func setUpRoutes(r chi.Router) {
	r.Group(routes.GetRoutes())
}
