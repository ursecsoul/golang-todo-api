package routes

import (
	"net/http"
	"todo-api/handlers"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /todos", handlers.CreateTodo)
	mux.HandleFunc("GET /todos", handlers.GetTodos)
	mux.HandleFunc("GET /todos/{id}", handlers.GetTodo)
	mux.HandleFunc("PUT /todos/{id}", handlers.UpdateTodo)
	mux.HandleFunc("DELETE /todos/{id}", handlers.DeleteTodo)

	return mux
}
