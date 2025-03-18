package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tjingsheng/jsask-forum-backend/internal/router"
)

func main() {
	r := router.Setup()
	fmt.Print("Server is running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
