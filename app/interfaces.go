package app

import (
	"net/http"
	"todomodule/domain"
)

type Repo interface {
	GetAll() ([]domain.Todo, error)
	Create(todo domain.Todo) error
	Update(todoNew domain.Todo, todoOld domain.Todo) error
	Delete(todo domain.Todo) error
}

type HttpRepo interface {
	GetAll(w http.ResponseWriter, r *http.Request) error
	Create()
}
