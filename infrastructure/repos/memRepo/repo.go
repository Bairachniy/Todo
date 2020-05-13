package memRepo

import (
	"todomodule/app"
	"todomodule/domain"
)

type memRepo struct {
	todo []domain.Todo
}

func (m memRepo) GetAll() ([]domain.Todo, error) {
	panic("implement me")
}

func (m memRepo) Create(todo domain.Todo) error {
	panic("implement me")
}

func NewMemRepo() app.Repo {
	return &memRepo{}
}
