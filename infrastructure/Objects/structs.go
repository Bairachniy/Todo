package Objects

import "todomodule/domain"

type TodoObj struct {
	todo []domain.Todo
}

type TodoTmp struct {
	id   int
	todo string
}
