package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tjingsheng/jsask-forum-backend/internal/database"
	"github.com/tjingsheng/jsask-forum-backend/internal/router"
)

func main() {
	Database := database.Init()
	// h := handlers.New(Database)

	fmt.Println(Database)

	r := router.Setup()
	fmt.Print("Listening on port 8000 at http://localhost:8000 !")
	log.Fatalln(http.ListenAndServe(":8000", r))
}
