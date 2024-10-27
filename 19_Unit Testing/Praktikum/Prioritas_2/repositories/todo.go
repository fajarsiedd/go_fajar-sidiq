package repositories

import (
	"go-todos-api/database"
	"go-todos-api/models"
)

type TodoRepositoryImpl struct {
}

func InitTodoRepository() TodoRepository {
	return &TodoRepositoryImpl{}
}

func (cr *TodoRepositoryImpl) GetAll() ([]models.Todo, error) {
	var todos []models.Todo

	if err := database.DB.Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}

func (cr *TodoRepositoryImpl) GetByID(id string) (models.Todo, error) {
	var todo models.Todo

	if err := database.DB.First(&todo, "id = ?", id).Error; err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func (cr *TodoRepositoryImpl) Create(todoInput models.TodoInput) (models.Todo, error) {
	var todo models.Todo = models.Todo{
		Title:       todoInput.Title,
		Description: todoInput.Description,
	}

	result := database.DB.Create(&todo)

	if err := result.Error; err != nil {
		return models.Todo{}, err
	}

	if err := result.Last(&todo).Error; err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func (cr *TodoRepositoryImpl) Update(todoInput models.TodoInput, id string) (models.Todo, error) {
	todo, err := cr.GetByID(id)

	if err != nil {
		return models.Todo{}, err
	}

	todo.Title = todoInput.Title
	todo.Description = todoInput.Description

	if err := database.DB.Save(&todo).Error; err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func (cr *TodoRepositoryImpl) Delete(id string) error {
	todo, err := cr.GetByID(id)

	if err != nil {
		return err
	}

	if err := database.DB.Unscoped().Delete(&todo).Error; err != nil {
		return err
	}

	return nil
}
