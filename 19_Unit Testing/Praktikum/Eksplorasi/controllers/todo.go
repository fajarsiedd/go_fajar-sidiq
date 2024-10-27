package controllers

import (
	"go-todos-api/models"
	"go-todos-api/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TodoController struct {
	Service services.TodoService
}

func InitTodoController() TodoController {
	return TodoController{
		Service: services.InitTodoService(),
	}
}

func (cc *TodoController) GetAll(c echo.Context) error {
	todos, err := cc.Service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse[string]{
			Status:  "failed",
			Message: "failed to fetch todo data",
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse[[]models.Todo]{
		Status:  "success",
		Message: "all todos",
		Data:    todos,
	})
}

func (cc *TodoController) GetByID(c echo.Context) error {
	todoID := c.Param("id")

	todo, err := cc.Service.GetByID(todoID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.BaseResponse[string]{
			Status:  "failed",
			Message: "todo not found",
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse[models.Todo]{
		Status:  "success",
		Message: "todo found",
		Data:    todo,
	})
}

func (cc *TodoController) Create(c echo.Context) error {
	var todoInput models.TodoInput

	if err := c.Bind(&todoInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: "request invalid",
		})
	}

	todo, err := cc.Service.Create(todoInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse[string]{
			Status:  "failed",
			Message: "failed to create a todo",
		})
	}

	return c.JSON(http.StatusCreated, models.BaseResponse[models.Todo]{
		Status:  "success",
		Message: "todo created",
		Data:    todo,
	})
}

func (cc *TodoController) Update(c echo.Context) error {
	todoID := c.Param("id")

	var todoInput models.TodoInput

	if err := c.Bind(&todoInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.BaseResponse[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	todo, err := cc.Service.Update(todoInput, todoID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse[string]{
			Status:  "failed",
			Message: "failed to update todo",
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse[models.Todo]{
		Status:  "success",
		Message: "todo updated",
		Data:    todo,
	})
}

func (cc *TodoController) Delete(c echo.Context) error {
	todoID := c.Param("id")

	err := cc.Service.Delete(todoID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse[string]{
			Status:  "failed",
			Message: "failed to delete todo",
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse[string]{
		Status:  "success",
		Message: "todo deleted",
	})
}
