package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-api/models"
	"todo-api/storage"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)

	if todo.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	todo.Status = "pending"
	created := storage.CreateTodo(todo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storage.GetAllTodos())
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	todo, found := storage.GetTodoByID(id)

	if !found {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))

	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)

	updated, found := storage.UpdateTodo(id, todo)
	if !found {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))

	if !storage.DeleteTodo(id) {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
