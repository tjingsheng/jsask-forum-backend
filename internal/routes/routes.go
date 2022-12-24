package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tjingsheng/jsask-forum-backend/internal/handlers"
)

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/posts", func(w http.ResponseWriter, req *http.Request) {
			response, _ := handlers.HandleListPosts(w, req)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})

		r.Get("/users", func(w http.ResponseWriter, req *http.Request) {
			response, _ := handlers.HandleListUsers(w, req)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})

		r.Get("/tags", func(w http.ResponseWriter, req *http.Request) {
			response, _ := handlers.HandleListTags(w, req)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})
	}
}
