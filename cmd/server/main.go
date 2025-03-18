package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tjingsheng/jsask-forum-backend/internal/routes"
)

func main() {
	fmt.Print("Server is running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", routes.Router()))
}
