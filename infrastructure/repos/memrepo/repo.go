package memrepo

import (
	"sync"
	"todomodule/app"
	"todomodule/domain"
)

type todoMap = map[string]domain.Todo

type memRepo struct {
	todo todoMap
	m    sync.Mutex
}

func (m *memRepo) Update(todo domain.Todo) error {
	m.m.Lock()
	m.todo[todo.ID()] = todo
	m.m.Unlock()
	return nil
}

func (m *memRepo) Delete(id string) error {
	m.m.Lock()
	delete(m.todo, id)
	m.m.Unlock()
	return nil
}

func (m *memRepo) GetAll() ([]domain.Todo, error) {
	todoTmp := make([]domain.Todo, 0, len(m.todo))
	m.m.Lock()
	for k := range m.todo {
		todoTmp = append(todoTmp, m.todo[k])
	}
	m.m.Unlock()
	return todoTmp, nil
}

func (m *memRepo) Create(todo domain.Todo) error {
	m.m.Lock()
	m.todo[todo.ID()] = todo
	m.m.Unlock()
	return nil
}

func NewMemRepo() app.Repo {
	return &memRepo{todo: make(todoMap)}
}
