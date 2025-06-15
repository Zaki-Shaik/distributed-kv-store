package kvstore

import (
	"errors"
	"sync"
)

type Store struct {
	mu    sync.RWMutex
	store map[string]string
}

func NewStore() *Store {
	return &Store{
		store: make(map[string]string),
	}
}

func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.store[key]
	return value, ok
}

func (s *Store) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if key == "" {
		return errors.New("key-cannot-be-empty")
	}

	delete(s.store, key)

	return nil
}

func (s *Store) Put(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[key] = value
}
