package tests

import (
	"errors"
	"go-todos-api/controllers"
	"go-todos-api/models"
	"go-todos-api/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	todoRepository DummyTodoRepository
	todoController controllers.TodoController
	todoService    services.TodoService
)

var defaultTime time.Time = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)

type DummyTodoRepository struct {
	ReturnTodos []models.Todo
	ReturnTodo  models.Todo
	ReturnError error
}

func (dummy *DummyTodoRepository) GetAll() ([]models.Todo, error) {
	return dummy.ReturnTodos, dummy.ReturnError
}

func (dummy *DummyTodoRepository) GetByID(id string) (models.Todo, error) {
	return dummy.ReturnTodo, dummy.ReturnError
}

func (dummy *DummyTodoRepository) Create(categoryInput models.TodoInput) (models.Todo, error) {
	return dummy.ReturnTodo, dummy.ReturnError
}

func (dummy *DummyTodoRepository) Update(categoryInput models.TodoInput, id string) (models.Todo, error) {
	return dummy.ReturnTodo, dummy.ReturnError
}

func (dummy *DummyTodoRepository) Delete(id string) error {
	return dummy.ReturnError
}

func setupTest() {
	todoRepository = DummyTodoRepository{}

	todoService = services.TodoService{Repository: &todoRepository}

	todoController = controllers.TodoController{Service: todoService}
}

func TestGetAll(t *testing.T) {
	setupTest()

	e := echo.New()

	var (
		req *http.Request
		rec *httptest.ResponseRecorder
		c   echo.Context
	)

	t.Run("get all todo success", func(t *testing.T) {
		req = httptest.NewRequest(http.MethodGet, "/todos", nil)
		rec = httptest.NewRecorder()
		c = e.NewContext(req, rec)

		todoRepository.ReturnTodos = []models.Todo{
			{
				ID:          1,
				Title:       "Task 1",
				Description: "Menyapu halaman",
				CreatedAt:   defaultTime,
				UpdatedAt:   defaultTime,
				DeletedAt:   gorm.DeletedAt{Time: defaultTime, Valid: false},
			},
		}

		todoController.GetAll(c)
		assert.Equal(t, http.StatusOK, rec.Code)

		expected := `{"status":"success","message":"all todos","data":[{"id":1,"title":"Task 1","description":"Menyapu halaman","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","deleted_at":null}]}`
		assert.JSONEq(t, expected, rec.Body.String())
	})

	t.Run("get all todo failed", func(t *testing.T) {
		req = httptest.NewRequest(http.MethodGet, "/todos", nil)
		rec = httptest.NewRecorder()
		c = e.NewContext(req, rec)

		todoRepository.ReturnError = errors.New("database error")

		todoController.GetAll(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		expected := `{"status":"failed","message":"failed to fetch todo data"}`
		assert.JSONEq(t, expected, rec.Body.String())
	})
}

func TestGetByID(t *testing.T) {
	setupTest()

	e := echo.New()

	var (
		req *http.Request
		rec *httptest.ResponseRecorder
		c   echo.Context
	)

	t.Run("get todo by id success", func(t *testing.T) {
		req = httptest.NewRequest(http.MethodGet, "/todos/1", nil)
		rec = httptest.NewRecorder()
		c = e.NewContext(req, rec)

		todoRepository.ReturnTodo = models.Todo{
			ID:          1,
			Title:       "Task 1",
			Description: "Menyapu halaman",
			CreatedAt:   defaultTime,
			UpdatedAt:   defaultTime,
			DeletedAt:   gorm.DeletedAt{Time: defaultTime, Valid: false},
		}

		todoController.GetByID(c)
		assert.Equal(t, http.StatusOK, rec.Code)

		expected := `{"status":"success","message":"todo found","data":{"id":1,"title":"Task 1","description":"Menyapu halaman","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","deleted_at":null}}`
		assert.JSONEq(t, expected, rec.Body.String())
	})

	t.Run("get todo by id failed", func(t *testing.T) {
		req = httptest.NewRequest(http.MethodGet, "/todos/0", nil)
		rec = httptest.NewRecorder()
		c = e.NewContext(req, rec)

		todoRepository.ReturnError = errors.New("record not found")

		todoController.GetByID(c)
		assert.Equal(t, http.StatusNotFound, rec.Code)

		expected := `{"status":"failed","message":"todo not found"}`
		assert.JSONEq(t, expected, rec.Body.String())
	})
}

func TestCreate(t *testing.T) {
	setupTest()

	e := echo.New()

	var (
		req *http.Request
		rec *httptest.ResponseRecorder
		c   echo.Context
	)

	t.Run("create todo success", func(t *testing.T) {
		todoInput := `{"title":"New Task","description":"Menyapu halaman"}`
		req = httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(todoInput))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec = httptest.NewRecorder()
		c = e.NewContext(req, rec)

		todoRepository.ReturnTodo = models.Todo{
			ID:          1,
			Title:       "New Task",
			Description: "Menyapu halaman",
			CreatedAt:   defaultTime,
			UpdatedAt:   defaultTime,
			DeletedAt:   gorm.DeletedAt{Time: defaultTime, Valid: false},
		}

		todoController.Create(c)
		assert.Equal(t, http.StatusCreated, rec.Code)

		expected := `{"status":"success","message":"todo created","data":{"id":1,"title":"New Task","description":"Menyapu halaman","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","deleted_at":null}}`
		assert.JSONEq(t, expected, rec.Body.String())
	})

	t.Run("create todo failed; bind error", func(t *testing.T) {
		todoInput := `{"title":"New Task","description":false}`
		req = httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(todoInput))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec = httptest.NewRecorder()
		c = e.NewContext(req, rec)

		todoController.Create(c)
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		expected := `{"status":"failed","message":"request invalid"}`
		assert.JSONEq(t, expected, rec.Body.String())
	})

	t.Run("create todo failed; internal server error", func(t *testing.T) {
		todoInput := `{"title":"New Task","description":"Menyapu halaman"}`
		req = httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(todoInput))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec = httptest.NewRecorder()
		c = e.NewContext(req, rec)

		todoRepository.ReturnError = errors.New("database error")

		todoController.Create(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		expected := `{"status":"failed","message":"failed to create a todo"}`
		assert.JSONEq(t, expected, rec.Body.String())
	})
}

func TestUpdate(t *testing.T) {
	setupTest()

	e := echo.New()

	var (
		req *http.Request
		rec *httptest.ResponseRecorder
		c   echo.Context
	)

	t.Run("update todo success", func(t *testing.T) {
		todoInput := `{"title":"New Task","description":"Menggosok pintu"}`
		req = httptest.NewRequest(http.MethodPost, "/todos/1", strings.NewReader(todoInput))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec = httptest.NewRecorder()
		c = e.NewContext(req, rec)

		todoRepository.ReturnTodo = models.Todo{
			ID:          1,
			Title:       "New Task",
			Description: "Menggosok pintu",
			CreatedAt:   defaultTime,
			UpdatedAt:   defaultTime,
			DeletedAt:   gorm.DeletedAt{Time: defaultTime, Valid: false},
		}

		todoController.Update(c)
		assert.Equal(t, http.StatusOK, rec.Code)

		expected := `{"status":"success","message":"todo updated","data":{"id":1,"title":"New Task","description":"Menggosok pintu","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","deleted_at":null}}`
		assert.JSONEq(t, expected, rec.Body.String())
	})

	t.Run("update todo failed; bind error", func(t *testing.T) {
		todoInput := `{"title":"New Task","description":false}`
		req = httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(todoInput))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec = httptest.NewRecorder()
		c = e.NewContext(req, rec)

		todoController.Update(c)
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		expected := `{"status":"failed","message":"invalid request"}`
		assert.JSONEq(t, expected, rec.Body.String())
	})

	t.Run("update todo failed; internal server error", func(t *testing.T) {
		todoInput := `{"title":"New Task","description":"Menggosok pintu"}`
		req = httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(todoInput))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec = httptest.NewRecorder()
		c = e.NewContext(req, rec)

		todoRepository.ReturnError = errors.New("database error")

		todoController.Update(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		expected := `{"status":"failed","message":"failed to update todo"}`
		assert.JSONEq(t, expected, rec.Body.String())
	})
}

func TestDelete(t *testing.T) {
	setupTest()

	e := echo.New()

	var (
		req *http.Request
		rec *httptest.ResponseRecorder
		c   echo.Context
	)

	t.Run("delete todo success", func(t *testing.T) {
		req = httptest.NewRequest(http.MethodGet, "/todos/1", nil)
		rec = httptest.NewRecorder()
		c = e.NewContext(req, rec)

		todoRepository.ReturnTodo = models.Todo{
			ID:          1,
			Title:       "Task 1",
			Description: "Menyapu halaman",
			CreatedAt:   defaultTime,
			UpdatedAt:   defaultTime,
			DeletedAt:   gorm.DeletedAt{Time: defaultTime, Valid: false},
		}

		todoController.Delete(c)
		assert.Equal(t, http.StatusOK, rec.Code)

		expected := `{"status":"success","message":"todo deleted"}`
		assert.JSONEq(t, expected, rec.Body.String())
	})

	t.Run("delete todo failed", func(t *testing.T) {
		req = httptest.NewRequest(http.MethodGet, "/todos/0", nil)
		rec = httptest.NewRecorder()
		c = e.NewContext(req, rec)

		todoRepository.ReturnError = errors.New("record not found")

		todoController.Delete(c)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		expected := `{"status":"failed","message":"failed to delete todo"}`
		assert.JSONEq(t, expected, rec.Body.String())
	})
}
