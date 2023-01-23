package routes

import (
	"github.com/go-chi/chi"
	"github.com/tjingsheng/jsask-forum-backend/internal/handlers"
)

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/tags", handlers.GetAllTags)
		r.Get("/curruser/{username}", handlers.GetSalt)
		r.Get("/curruser/{username}/{password}", handlers.GetCurrUser)
		r.Get("/posts/{userId}", handlers.GetAllPosts)
		r.Get("/currpost/{userId}/{postId}", handlers.GetCurrPost)

		r.Put("/postpreference", handlers.PutPostPreference)
		r.Put("/post", handlers.PutPost)

		r.Post("/post", handlers.PostPost)
		r.Post("/curruser", handlers.PostCurrUser)

		r.Delete("/post/{postId}", handlers.DeletePost)

	}
}
