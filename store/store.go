package store

import (
	"sync"
)

type Store[T any] struct {
	m *sync.Mutex
	d map[string]T
}

func NewStore[T any]() *Store[T] {
	return &Store[T]{
		m: &sync.Mutex{},
		d: map[string]T{},
	}
}

func (s *Store[T]) Insert(key string, t T) bool {
	s.m.Lock()
	defer s.m.Unlock()

	if _, ok := s.d[key]; ok {
		return false
	}
	s.d[key] = t
	return true
}
