package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tjingsheng/jsask-forum-backend/internal/retriever"
)

func GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/users", func(w http.ResponseWriter, req *http.Request) {
			response, _ := retriever.HandleList(w, req, "users")

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		})
	}
}
