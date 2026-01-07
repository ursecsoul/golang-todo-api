package storage

import (
	"sync"
	"time"
	"todo-api/models"
)

var (
	Todos  = []models.Todo{}
	lastID = 0
	mutex  sync.Mutex
)

func CreateTodo(todo models.Todo) models.Todo {
	mutex.Lock()
	defer mutex.Unlock()

	lastID++
	todo.ID = lastID
	todo.CreatedAt = time.Now()
	Todos = append(Todos, todo)
	return todo
}

func GetAllTodos() []models.Todo {
	return Todos
}

func GetTodoByID(id int) (models.Todo, bool) {
	for _, todo := range Todos {
		if todo.ID == id {
			return todo, true
		}
	}
	return models.Todo{}, false
}

func UpdateTodo(id int, updated models.Todo) (models.Todo, bool) {
	for i, todo := range Todos {
		if todo.ID == id {
			Todos[i].Title = updated.Title
			Todos[i].Description = updated.Description
			Todos[i].Status = updated.Status
			return Todos[i], true
		}
	}
	return models.Todo{}, false
}

func DeleteTodo(id int) bool {
	for i, todo := range Todos {
		if todo.ID == id {
			Todos = append(Todos[:i], Todos[i+1:]...)
			return true
		}
	}
	return false
}
