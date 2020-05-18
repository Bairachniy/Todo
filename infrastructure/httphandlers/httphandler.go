package httphandlers

import (
	"bytes"
	"encoding/json"
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

func NewHttpHandler(todo *app.Todos) httpHandler {
	httpHandler := httpHandler{todo: todo}
	return httpHandler
}

func (h *httpHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo httpTodo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.todo.Create(domain.Todo{
		Name: todo.Name,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h httpHandler) GetAllTodo(w http.ResponseWriter, r *http.Request) {
	all, err := h.todo.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var b bytes.Buffer
	err = json.NewEncoder(&b).Encode(all)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b.Bytes())
}
