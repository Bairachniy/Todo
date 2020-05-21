package memrepo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMemRepo(t *testing.T) {
	memRepo := NewMemRepo()
	assert.NotNil(t, memRepo)
}

func TestMemRepo_Create(t *testing.T) {

}
