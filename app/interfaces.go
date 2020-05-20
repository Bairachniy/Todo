package app

import (
	"todomodule/domain"
)

type TodoService interface {
	Create(todoCmd CreateTodoCommand) error
	Update(todoCmd UpdateTodoCommand) error
	Delete(id string) error
	GetAll() ([]domain.Todo, error)
}

type Repo interface {
	Create(todo domain.Todo) error
	Update(todo domain.Todo) error
	GetAll() ([]domain.Todo, error)
	Delete(id string) error
}
