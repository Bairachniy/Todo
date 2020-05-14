package memRepo

import (
	"fmt"
	"sync"
	"todomodule/app"
	"todomodule/domain"
)

type todoMap = map[int]domain.Todo

type memRepo struct {
	todo todoMap
	m    sync.Mutex
}

func (m memRepo) GetAll() ([]domain.Todo, error) {
	todoTmp := []domain.Todo{}
	m.m.Lock()
	for k := range m.todo {
		todoTmp = append(todoTmp, domain.Todo{m.todo[k].Name})
	}
	m.m.Unlock()
	return todoTmp, nil
}

func (m memRepo) Create(todo domain.Todo) error {
	m.m.Lock()
	var keys []int
	for k := range m.todo {
		keys = append(keys, k)
	}
	fmt.Println(len(keys), "maps")
	m.todo[len(keys)+1] = todo
	m.m.Unlock()
	return nil
}

func NewMemRepo() app.Repo {
	return &memRepo{todo: make(todoMap)}
}
