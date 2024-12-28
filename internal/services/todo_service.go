package services

import "errors"

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos []Todo
var nextID = 1

func GetAllTodos() *[]Todo {
	return &todos
}

func CreateTodo(title string) *Todo {
	newTodo := Todo{
		ID:        nextID,
		Title:     title,
		Completed: false,
	}
	nextID++
	todos = append(todos, newTodo)
	return &newTodo
}

func UpdateTodoByID(id int, title string, completed bool) (*Todo, error) {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Title = title
			todos[i].Completed = completed
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func DeleteTodoByID(id int) error {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}
