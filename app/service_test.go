package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTodos(t *testing.T) {
	var repo Repo
	createdTodo := NewTodos(repo)
	assert.NotNil(t, createdTodo)
}
