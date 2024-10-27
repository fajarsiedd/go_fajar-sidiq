package repositories

import "go-todos-api/models"

type TodoRepository interface {
	GetAll() ([]models.Todo, error)
	GetByID(id string) (models.Todo, error)
	Create(categoryInput models.TodoInput) (models.Todo, error)
	Update(categoryInput models.TodoInput, id string) (models.Todo, error)
	Delete(id string) error
}
