package store

import (
	"sync"
)

type Store[T any] struct {
	m *sync.RWMutex
	d map[string]T
}

func NewStore[T any]() *Store[T] {
	return &Store[T]{
		m: &sync.RWMutex{},
		d: map[string]T{},
	}
}

func (s *Store[T]) Get(key string) (T, bool) {
	s.m.Lock()
	defer s.m.Unlock()

	t, ok := s.d[key]
	return t, ok
}

func (s *Store[T]) Insert(key string, t T) bool {
	s.m.Lock()
	defer s.m.Unlock()

	// if _, ok := s.d[key]; ok {
	// 	return false
	// }
	s.d[key] = t
	return true
}

func (s *Store[T]) Range(f func(id string)) {
	// s.m.RLock()
	// defer s.m.RUnlock()
	for key := range s.d {
		f(key)
	}
}
