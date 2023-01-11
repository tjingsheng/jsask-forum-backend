package routes

import (
	"github.com/go-chi/chi"
	"github.com/tjingsheng/jsask-forum-backend/internal/handlers"
)

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/tags", handlers.GetAllTags)
		r.Get("/curruser/{username}", handlers.GetCurrUser)
		r.Get("/posts/{userId}", handlers.GetAllPosts)
		r.Get("/currpost/{userId}/{postId}", handlers.GetCurrPost)

		r.Post("/user", handlers.PostNewUser)
	}
}
