package httphandlers

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"todomodule/app"
	"todomodule/domain"
)

type HTTPHandler struct {
	todo app.TodoService
}

type httpTodo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewHTTPHandler(todo app.TodoService) HTTPHandler {
	httpHandler := HTTPHandler{todo: todo}
	return httpHandler
}

func (h *HTTPHandler) CreateTodo(c echo.Context) error {
	var todo httpTodo
	if err := json.NewDecoder(c.Request().Body).Decode(&todo); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	err := h.todo.Create(app.CreateTodoCommand{Name: todo.Name})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusCreated, todo)
}
func (h HTTPHandler) GetAllTodo(c echo.Context) error {
	todos, err := h.todo.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(todos)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, h.toHTTPTodo(todos))
}

func (h HTTPHandler) toHTTPTodo(todos []domain.Todo) []httpTodo {
	httpTodos := make([]httpTodo, 0, len(todos))
	for i := range todos {
		httpTodos = append(httpTodos, httpTodo{
			ID:   todos[i].ID(),
			Name: todos[i].Name(),
		})
	}
	return httpTodos
}

func (h HTTPHandler) DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	if err := h.todo.Delete(id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

func (h HTTPHandler) UpdateTodo(c echo.Context) error {
	var m app.UpdateTodoCommand
	m.ID = c.Param("id")
	if err := json.NewDecoder(c.Request().Body).Decode(&m); err != nil {
		return echo.NewHTTPError(http.StatusConflict)
	}
	if err := h.todo.Update(m); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, m)
}
