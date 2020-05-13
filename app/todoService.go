package app

import "todomodule/domain"

type Todos struct {
	repo Repo
}

func NewTodos(repo Repo) *Todos {
	return &Todos{repo: repo}
}

func (t *Todos) Create(todo domain.Todo) error {
	return t.repo.Create(todo)
}
func (t *Todos) GetAll() ([]domain.Todo, error) {
	return t.repo.GetAll()
}
