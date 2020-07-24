package domain

import "testing"

func TestNewTodo(t *testing.T) {
	todo, err := NewTodo("Testing todo")
	if err != nil {
		panic(err)
	}
	expected := Todo{id: "", name: "Testing todo"}
	if todo.Name() != expected.Name() {
		t.Error("expected the same name in todo")
	}
}

func TestParseTodo(t *testing.T) {
	actual, err := ParseTodo("123", "Test todo")
	if err != nil {
		panic(err)
	}
	expected := Todo{
		id:   "123",
		name: "Test todo",
	}
	if actual != expected {
		t.Error("something wrong!!")
	}
}
