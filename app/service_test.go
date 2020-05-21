package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"todomodule/infrastructure/repos/memRepo"
)

func TestNewMemRepo(t *testing.T) {
	repo := memRepo.NewMemRepo()
	assert.NotNil(t, repo)
}

func TestNewTodos(t *testing.T) {
	var repo Repo
	createdTodo := NewTodos(repo)
	assert.NotNil(t, createdTodo)
}
