package httphandlers

import (
	"fmt"
	"net/http"
	"todomodule/app"
	"todomodule/domain"
)

type httpHandler struct {
	todo *app.Todos
}

func NewHttpHandler(todo *app.Todos) httpHandler {
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		todo.Create(domain.Todo{r.URL.Query().Get("todo")})
	})
	http.HandleFunc("/getall", func(w http.ResponseWriter, r *http.Request) {
		all, err := todo.GetAll()
		if err != nil {
			panic(err)
		}
		for value, index := range all {
			fmt.Fprintf(w, "%s,%s", index, value)
		}
	})
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
	return httpHandler{todo: todo}
}

func (h httpHandler) GetAll() ([]domain.Todo, error) {
	todo, err := h.todo.GetAll()
	return todo, err
}

func (h httpHandler) Create(todo domain.Todo) error {
	return h.todo.Create(todo)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

//func CreateTodo(w http.ResponseWriter, r *http.Request, todo *){
//
//}
//
//func (h httpHandler)GetAllTodo(w http.ResponseWriter, r *http.Request)  {
//
//}
