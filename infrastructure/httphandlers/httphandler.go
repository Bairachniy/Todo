package httphandlers

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"todomodule/app"
	"todomodule/domain"
)

type httpHandler struct {
	todo *app.Todos
}

type httpTodo struct {
	Name string `json:"name"`
}
type udateTodo struct {
	OldName string `json:"old"`
	NewName string `json:"new"`
}

func NewHttpHandler(todo *app.Todos) httpHandler {
	httpHandler := httpHandler{todo: todo}
	return httpHandler
}

//func (h *httpHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
//	var todo httpTodo
//	err := json.NewDecoder(r.Body).Decode(&todo)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	err = h.todo.Create(domain.Todo{
//		Name: todo.Name,
//	})
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	w.WriteHeader(http.StatusOK)
//}

//func (h httpHandler) GetAllTodo(w http.ResponseWriter, r *http.Request) {
//	all, err := h.todo.GetAll()
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	var b bytes.Buffer
//	err = json.NewEncoder(&b).Encode(all)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	w.WriteHeader(http.StatusOK)
//	w.Write(b.Bytes())
//}
func (h *httpHandler) CreateTodo(c echo.Context) error {
	var todo httpTodo
	if err := json.NewDecoder(c.Request().Body).Decode(&todo); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	err := h.todo.Create(domain.Todo{
		Name: todo.Name,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusCreated, todo)
}
func (h httpHandler) GetAllTodo(c echo.Context) error {
	all, err := h.todo.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(all)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, all)
}
func (h httpHandler) DeleteTodo(c echo.Context) error {
	var todo httpTodo
	if err := json.NewDecoder(c.Request().Body).Decode(&todo); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	if err := h.todo.Delete(domain.Todo{
		Name: todo.Name,
	}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, todo)
}

//{
//	"update":{
//		"old":"Tratata",
//		"new":"Pamparirurai"
//	}
//}
func (h httpHandler) UpdateTodo(c echo.Context) error {
	m := make(map[string]udateTodo)
	if err := json.NewDecoder(c.Request().Body).Decode(&m); err != nil {
		return echo.NewHTTPError(http.StatusConflict)
	}
	if err := h.todo.Update(domain.Todo{
		Name: m["update"].NewName,
	}, domain.Todo{Name: m["update"].OldName}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, m)
}
