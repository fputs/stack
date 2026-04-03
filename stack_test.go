package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	s := New[int]()
	assert.NotNil(t, s)
}

func TestStackPush(t *testing.T) {
	s := New[int]()
	s.Push(2)
	s.Push(10)

	p, err := s.Peek()
	assert.Equal(t, p, 10)
	assert.NoError(t, err)

	pp, err := s.PeekPtr()
	assert.NotNil(t, pp)
	assert.NoError(t, err)
	assert.Equal(t, *pp, 10)
}

func TestStackPop(t *testing.T) {
	s := New[int](2, 10)

	x, err := s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, x, 10)

	y, err := s.PopPtr()
	assert.NotNil(t, y)
	assert.NoError(t, err)
	assert.Equal(t, *y, 2)
}

func TestStackPopEmpty(t *testing.T) {
	s := New[int]()
	x, err := s.Pop()
	assert.Error(t, err)
	assert.Zero(t, x)
}

func TestStackPopPtrEmpty(t *testing.T) {
	s := New[int]()
	x, err := s.PopPtr()
	assert.Error(t, err)
	assert.Nil(t, x)
}

func TestStackClear(t *testing.T) {
	s := New[int]()
	s.Push(10)
	s.Push(11)
	assert.Equal(t, s.Size(), 2)
	s.Clear()
	assert.Equal(t, s.Size(), 0)
}

func TestStackComplete(t *testing.T) {
	s := New[int](1, 2)
	s.Push(3)

	assert.Equal(t, s.Size(), 3)
	x, err := s.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 3, x)

	y, err := s.PeekPtr()
	assert.NotNil(t, y)
	assert.NoError(t, err)
	assert.Equal(t, 3, *y)

	z, err := s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, z, 3)

	p, err := s.Peek()
	assert.NoError(t, err)
	assert.Equal(t, p, 2)
	assert.Equal(t, s.Size(), 2)

	s.Clear()
	assert.Equal(t, s.Size(), 0)
}
