package stack

import "fmt"

type Stack[T any] struct {
	arr  []T
	size int
}

func New[T any](values ...T) *Stack[T] {
	if len(values) > 0 {
		return &Stack[T]{
			arr:  values,
			size: len(values),
		}
	}

	return &Stack[T]{
		arr:  make([]T, 0),
		size: 0,
	}
}

func (s *Stack[T]) Clear() {
	s.arr = make([]T, 0)
	s.size = 0
}

func (s *Stack[T]) Size() int {
	return s.size
}

func (s *Stack[T]) Pop() (T, error) {
	if s.size == 0 {
		var t T
		return t, fmt.Errorf("tried to pop empty stack")
	}
	t := s.arr[s.size-1]
	s.arr = s.arr[:s.size-1]
	s.size--
	return t, nil
}

func (s *Stack[T]) PopPtr() (*T, error) {
	if s.size == 0 {
		return nil, fmt.Errorf("tried to pop empty stack")
	}
	t := s.arr[s.size-1]
	s.arr = s.arr[:s.size-1]
	s.size--
	return &t, nil
}

func (s *Stack[T]) Push(t T) {
	s.arr = append(s.arr, t)
	s.size++
}

func (s *Stack[T]) Peek() (T, error) {
	if s.size == 0 {
		var t T
		return t, fmt.Errorf("tried to peek empty stack")
	}
	t := s.arr[s.size-1]
	return t, nil
}

func (s *Stack[T]) PeekPtr() (*T, error) {
	if s.size == 0 {
		return nil, fmt.Errorf("tried to peek empty stack")
	}
	return &(s.arr[s.size-1]), nil
}
