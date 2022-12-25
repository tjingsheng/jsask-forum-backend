package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tjingsheng/jsask-forum-backend/internal/handlers"
)

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/comments", func(w http.ResponseWriter, req *http.Request) {
			response, _ := handlers.HandleListComments(w, req)

			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			json.NewEncoder(w).Encode(response)
		})

		r.Get("/posts", func(w http.ResponseWriter, req *http.Request) {
			response, _ := handlers.HandleListPosts(w, req)

			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			json.NewEncoder(w).Encode(response)
		})

		r.Get("/tags", func(w http.ResponseWriter, req *http.Request) {
			response, _ := handlers.HandleListTags(w, req)

			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			json.NewEncoder(w).Encode(response)
		})

		r.Get("/curruser", func(w http.ResponseWriter, req *http.Request) {
			response, _ := handlers.HandleCurrUser(w, req)

			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			json.NewEncoder(w).Encode(response)
		})
	}
}
