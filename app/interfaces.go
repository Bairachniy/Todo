package app

import "todomodule/domain"

type Repo interface {
	GetAll() ([]domain.Todo, error)
	Create(todo domain.Todo) error
}
