package domain

import "errors"
import uuid "github.com/satori/go.uuid"

var ErrInvalidTodoName = errors.New("name is invalid")

type Todo struct {
	id   string
	name string
}

func ParseTodo(id, name string) (Todo, error) {
	if len(name) == 0 || len(name) > 100 {
		return Todo{}, ErrInvalidTodoName
	}
	return Todo{
		id:   id,
		name: name,
	}, nil
}

func NewTodo(name string) (Todo, error) {
	if len(name) == 0 || len(name) > 100 {
		return Todo{}, ErrInvalidTodoName
	}
	return Todo{
		id:   uuid.NewV4().String(),
		name: name,
	}, nil
}

func (todo Todo) ID() string {
	return todo.id
}

func (todo Todo) Name() string {
	return todo.name
}
