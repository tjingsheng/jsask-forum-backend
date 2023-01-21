package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tjingsheng/jsask-forum-backend/internal/router"
)

func main() {
	port := os.Getenv("PORT")
	r := router.Setup()
	fmt.Print("Listening on: " + port)
	log.Fatal(http.ListenAndServe(":8000", r))
}
