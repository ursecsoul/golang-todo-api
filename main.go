package main

import (
	"log"
	"net/http"
	"todo-api/routes"
)

func main() {
	mux := routes.RegisterRoutes()

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
