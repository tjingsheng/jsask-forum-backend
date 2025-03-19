package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/routes"
)

func main() {
	database.InitDB()
	port := "8000"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = val
	}

	r := chi.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Received request: %s %s", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	})

	log.Printf("Server is running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, routes.Router()))
}
