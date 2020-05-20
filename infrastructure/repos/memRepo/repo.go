package memRepo

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

func (m memRepo) Update(todoNew domain.Todo, todoOld domain.Todo) error {
	m.Delete(todoOld)
	m.m.Lock()
	m.todo[todoNew.Name] = todoNew
	m.m.Unlock()
	return nil
}

func (m memRepo) Delete(todo domain.Todo) error {
	m.m.Lock()
	delete(m.todo, todo.Name)
	m.m.Unlock()
	return nil
}

func (m memRepo) GetAll() ([]domain.Todo, error) {
	todoTmp := []domain.Todo{}
	m.m.Lock()
	for k := range m.todo {
		todoTmp = append(todoTmp, m.todo[k])
	}
	m.m.Unlock()
	return todoTmp, nil
}

func (m memRepo) Create(todo domain.Todo) error {
	m.m.Lock()
	m.todo[todo.Name] = todo
	m.m.Unlock()
	return nil
}

func NewMemRepo() app.Repo {
	return &memRepo{todo: make(todoMap)}
}
