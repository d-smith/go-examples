package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	s := new(Stack)
	assert.Equal(t, 0, s.Size())
	s.Push("a")
	s.Push("b")

	assert.Equal(t, 2, s.Size())

	assert.Equal(t, "b", s.Peek().(string))
	assert.Equal(t, "b", s.Pop().(string))
	assert.Equal(t, "a", s.Pop().(string))

	assert.Equal(t, 0, s.Size())
}
