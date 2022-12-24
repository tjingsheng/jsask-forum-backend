package router

import (
	"github.com/go-chi/chi"
	"github.com/tjingsheng/jsask-forum-backend/internal/routes"
)

func Setup() chi.Router {
	r := chi.NewRouter()
	setUpRoutes(r)
	return r
}

func setUpRoutes(r chi.Router) {
	r.Group(routes.GetRoutes())
}
