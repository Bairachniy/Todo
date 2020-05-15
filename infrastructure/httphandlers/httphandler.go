package httphandlers

import (
	"fmt"
	"log"
	"net/http"
	"todomodule/app"
	"todomodule/domain"
)

type httpHandler struct {
	todo *app.Todos
}

func NewHttpHandler(todo *app.Todos) httpHandler {
	http.HandleFunc("/", HelloServer)
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

//
//func (h Handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
//	fmt.Fprint(resp, h.Message)
//}

// функция добавления данных

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		//_________________________________________________________
		http.Redirect(w, r, "/", 301)
	} else {
		http.ServeFile(w, r, "templates/todo.html")
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

}
