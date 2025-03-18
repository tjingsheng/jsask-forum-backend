package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tjingsheng/jsask-forum-backend/internal/handlers"
)

func Router() chi.Router {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("healthy"))
	})

	r.Group(func(r chi.Router) {
		r.Get("/user/{username}/{password}", handlers.GetUser)
		r.Get("/user/{username}", handlers.GetUserSalt)
		r.Post("/user", handlers.PostUser)

		r.Get("/post/{userId}/{postId}", handlers.GetCurrPost)
		r.Get("/post/{userId}", handlers.GetAllPosts)
		r.Post("/post", handlers.PostPost)
		r.Put("/post", handlers.PutPost)
		r.Delete("/post/{postId}", handlers.DeletePost)

		r.Get("/tags", handlers.GetAllTags)

		r.Put("/postpreference", handlers.PutPostPreference)
	})
	return r
}
