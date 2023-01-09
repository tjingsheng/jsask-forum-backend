package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tjingsheng/jsask-forum-backend/internal/handlers"
)

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/tags", func(w http.ResponseWriter, req *http.Request) {
			response, _ := handlers.GetAllTags(w, req)
			json.NewEncoder(w).Encode(response)
		})

		r.Get("/curruser/{username}", func(w http.ResponseWriter, req *http.Request) {
			username := chi.URLParam(req, "username")
			response, _ := handlers.GetCurrentUser(w, req, username)
			json.NewEncoder(w).Encode(response)
		})

		r.Get("/posts", func(w http.ResponseWriter, req *http.Request) {
			response, _ := handlers.GetAllPosts(w, req)
			json.NewEncoder(w).Encode(response)
		})

	}
}
