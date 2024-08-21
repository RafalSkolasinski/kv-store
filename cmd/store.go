package main

import (
	"errors"
	"sync"
)

type Store struct {
	sync.RWMutex
	values map[string]string
}

var ErrNoKey = errors.New("key not in store")

func NewStore() *Store {
	return &Store{values: make(map[string]string)}
}

func (s *Store) Put(key, value string) error {
	s.Lock()
	s.values[key] = value
	s.Unlock()
	return nil
}

func (s *Store) Get(key string) (string, error) {
	s.RLock()
	value, ok := s.values[key]
	s.RUnlock()
	if !ok {
		return "", ErrNoKey
	}
	return value, nil
}

func (s *Store) Delete(key string) error {
	delete(s.values, key)
	return nil
}
