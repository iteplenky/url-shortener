package store

import "sync"

type Store struct {
	sync.RWMutex
	data map[string]string
}

func New() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Set(key, value string) {
	s.Lock()
	defer s.Unlock()
	s.data[key] = value
}

func (s *Store) Get(key string) (string, bool) {
	s.RLock()
	defer s.RUnlock()
	value, ok := s.data[key]
	return value, ok
}

func (s *Store) Delete(key string) {
	s.Lock()
	defer s.Unlock()
	delete(s.data, key)
}
