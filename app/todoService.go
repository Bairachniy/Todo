package app

import "todomodule/domain"

type Todos struct {
	repo Repo
}

type CreateTodoCommand struct {
	Name string
}

type UpdateTodoCommand struct {
	ID   string
	Name string
}

func NewTodos(repo Repo) *Todos {
	return &Todos{repo: repo}
}

func (t *Todos) Create(todoCmd CreateTodoCommand) error {
	todo, err := domain.NewTodo(todoCmd.Name)
	if err != nil {
		return err
	}
	return t.repo.Create(todo)
}

func (t *Todos) Update(todoCmd UpdateTodoCommand) error {
	todo, err := domain.ParseTodo(todoCmd.ID, todoCmd.Name)
	if err != nil {
		return err
	}
	return t.repo.Update(todo)
}

func (t *Todos) Delete(id string) error {
	return t.repo.Delete(id)
}

func (t *Todos) GetAll() ([]domain.Todo, error) {
	return t.repo.GetAll()
}
