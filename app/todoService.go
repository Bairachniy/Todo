package app

import "todomodule/domain"

type Todos struct {
	repo Repo
}

func (t *Todos) Update(todoNew domain.Todo, todoOld domain.Todo) error {
	return t.repo.Update(todoNew, todoOld)
}

func (t *Todos) Delete(todo domain.Todo) error {
	return t.repo.Delete(todo)
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
