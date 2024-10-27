package services

import (
	"go-todos-api/models"
	"go-todos-api/repositories"
)

type TodoService struct {
	Repository repositories.TodoRepository
}

func InitTodoService() TodoService {
	return TodoService{
		Repository: repositories.InitTodoRepository(),
	}
}

func (ts *TodoService) GetAll() ([]models.Todo, error) {
	return ts.Repository.GetAll()
}

func (ts *TodoService) GetByID(id string) (models.Todo, error) {
	return ts.Repository.GetByID(id)
}

func (ts *TodoService) Create(todoInput models.TodoInput) (models.Todo, error) {
	return ts.Repository.Create(todoInput)
}

func (ts *TodoService) Update(todoInput models.TodoInput, id string) (models.Todo, error) {
	return ts.Repository.Update(todoInput, id)
}

func (ts *TodoService) Delete(id string) error {
	return ts.Repository.Delete(id)
}
