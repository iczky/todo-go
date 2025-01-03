package services

import (
	"errors"
	"github.com/iczky/todo-fiber/internal/db"
	"github.com/iczky/todo-fiber/internal/models"
	"time"
)

type TodoService interface {
	GetAllTodos() (*[]models.Todo, error)
	GetTodoById(id int) (*models.Todo, error)
	CreateTodo(title string, description string, priority int, dueDate *string) (*models.Todo, error)
	UpdateTodoByID(id int, title string, description string, completed bool, priority int, dueDate *string) (*models.Todo, error)
	DeleteTodoByID(id int) error
}

type todoServiceImpl struct{}

func NewTodoService() TodoService {
	return &todoServiceImpl{}
}

func (s *todoServiceImpl) GetTodoById(id int) (*models.Todo, error) {
	var todo models.Todo
	result := db.DB.Find(&todo, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &todo, nil
}

func (s *todoServiceImpl) GetAllTodos() (*[]models.Todo, error) {
	var todos []models.Todo
	result := db.DB.Find(&todos) // Runs SELECT * FROM todos;
	if result.Error != nil {
		return nil, result.Error
	}
	return &todos, nil
}

func (s *todoServiceImpl) CreateTodo(title string, description string, priority int, dueDate *string) (*models.Todo, error) {
	todo := models.Todo{
		Title:       title,
		Description: description,
		Priority:    priority,
		Completed:   false,
	}

	if dueDate != nil {
		parsedDate, err := time.Parse("2006-01-02", *dueDate)
		if err != nil {
			return nil, errors.New("invalid due date format, use YYYY-MM-DD")
		}
		todo.DueDate = parsedDate
	}

	result := db.DB.Create(&todo)
	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}

func (s *todoServiceImpl) UpdateTodoByID(id int, title string, description string, completed bool, priority int, dueDate *string) (*models.Todo, error) {
	var todo models.Todo
	result := db.DB.First(&todo, id) // Runs SELECT * FROM todos WHERE id = ? (retrieve the first result)
	if result.Error != nil {
		return nil, errors.New("todo not found")
	}

	// Update fields
	todo.Title = title
	todo.Description = description
	todo.Completed = completed
	todo.Priority = priority

	// Parse and update the dueDate if necessary
	if dueDate != nil {
		parsedDate, err := time.Parse("2006-01-02", *dueDate) // Expecting YYYY-MM-DD
		if err != nil {
			return nil, errors.New("invalid due date format, use YYYY-MM-DD")
		}
		todo.DueDate = parsedDate
	}

	db.DB.Save(&todo) // Runs UPDATE todos SET ...

	return &todo, nil
}

func (s *todoServiceImpl) DeleteTodoByID(id int) error {
	result := db.DB.Delete(&models.Todo{}, id) // Runs DELETE FROM todos WHERE id = ?
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("todo not found")
	}
	return nil
}
