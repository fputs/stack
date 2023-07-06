package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack[int]()
	assert.NotNil(t, s)
}

func TestStackPush(t *testing.T) {
	s := NewStack[int]()
	s.Push(2)
	s.Push(10)

	assert.Equal(t, s.Peek(), 10)
	assert.Equal(t, *s.PeekPtr(), 10)
}

func TestStackPop(t *testing.T) {
	s := NewStack[int]()
	s.Push(2)
	s.Push(10)

	x, err := s.Pop()
	assert.NotNil(t, x)
	assert.NoError(t, err)
	assert.Equal(t, *x, 10)
}

func TestStackPopEmpty(t *testing.T) {
	s := NewStack[int]()
	x, err := s.Pop()
	assert.Error(t, err)
	assert.Nil(t, x)
}

func TestStackClear(t *testing.T) {
	s := NewStack[int]()
	s.Push(10)
	s.Push(11)
	assert.Equal(t, s.Size(), 2)
	s.Clear()
	assert.Equal(t, s.Size(), 0)
}

func TestStackComplete(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, s.Size(), 3)
	x := s.Peek()
	assert.Equal(t, 3, x)

	y := s.PeekPtr()
	assert.NotNil(t, y)
	assert.Equal(t, 3, *y)

	z, err := s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, *z, 3)
	assert.Equal(t, s.Peek(), 2)
	assert.Equal(t, s.Size(), 2)

	s.Clear()
	assert.Equal(t, s.Size(), 0)
}
