package memRepo

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"todomodule/app"
)

var repo app.Repo

func TestNewMemRepo(t *testing.T) {
	memRepo := NewMemRepo()
	assert.NotNil(t, memRepo)
}

func TestMemRepo_Create(t *testing.T) {

}
