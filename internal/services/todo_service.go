package services

import (
	"errors"
	"github.com/iczky/todo-fiber/internal/models"
)

type TodoService interface {
	GetAllTodos() *[]models.Todo
	CreateTodo(title string) *models.Todo
	UpdateTodoByID(id int, title string, completed bool) (*models.Todo, error)
	DeleteTodoByID(id int) error
}

type todoServiceImpl struct {
	todos  []models.Todo
	nextID int
}

func NewTodoService() TodoService {
	return &todoServiceImpl{
		todos:  make([]models.Todo, 0),
		nextID: 1,
	}
}

func (s *todoServiceImpl) GetAllTodos() *[]models.Todo {
	return &s.todos
}

func (s *todoServiceImpl) CreateTodo(title string) *models.Todo {
	newTodo := models.Todo{
		ID:        s.nextID,
		Title:     title,
		Completed: false,
	}
	s.nextID++
	s.todos = append(s.todos, newTodo)
	return &newTodo
}

func (s *todoServiceImpl) UpdateTodoByID(id int, title string, completed bool) (*models.Todo, error) {
	for i, todo := range s.todos {
		if todo.ID == id {
			s.todos[i].Title = title
			s.todos[i].Completed = completed
			return &s.todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func (s *todoServiceImpl) DeleteTodoByID(id int) error {
	for i, todo := range s.todos {
		if todo.ID == id {
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}
