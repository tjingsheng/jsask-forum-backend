package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tjingsheng/jsask-forum-backend/internal/router"
)

func main() {
	r := router.Setup()
	fmt.Print("Server is running on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
