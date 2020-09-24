package memory

import (
	"fmt"
	"sync"
)

type Storage struct {
	sync.RWMutex
	values map[string]interface{}
}

// New returns a new memory provider configured
func New() *Storage {
	return &Storage{
		values: make(map[string]interface{}),
	}
}

func (s *Storage) Get(key string) (interface{}, bool) {
	s.RLock()
	val, ok := s.values[key]
	s.RUnlock()
	return val, ok
}

func (s *Storage) Set(key string, val interface{}) {
	s.Lock()
	fmt.Println(s.values)
	s.values[key] = val
	s.Unlock()
}

func (s *Storage) Delete(key string) {
	s.RLock()
	_, ok := s.values[key]
	s.RUnlock()
	if ok {
		s.Lock()
		delete(s.values, key)
		s.Unlock()
	}
}
