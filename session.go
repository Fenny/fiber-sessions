package session

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

// Session represents a single session
type Session struct {
	sync.RWMutex
	ctx    *fiber.Ctx
	store  *Store
	values map[string]interface{}
	id     string
	new    bool
}

// New ...
func (s *Session) New() bool {
	return s.new
}

// ID ...
func (s *Session) ID() string {
	return s.id
}

// Set ...
func (s *Session) Set(key string, val interface{}) {
	s.Lock()
	s.values[key] = val
	s.Unlock()
}

// Get ...
func (s *Session) Get(key string) (interface{}, bool) {
	s.RLock()
	val, ok := s.values[key]
	s.RUnlock()
	return val, ok
}

// Delete ...
func (s *Session) Delete(key string) {
	s.RLock()
	_, ok := s.values[key]
	s.RUnlock()
	if ok {
		s.Lock()
		delete(s.values, key)
		s.Unlock()
	}
}

// Destroy ...
func (s *Session) Destroy() {
	s.Lock()
	s.values = make(map[string]interface{})
	s.Unlock()
}

// Save ...
func (s *Session) Save() error {
	s.RLock()
	values := s.values
	s.RUnlock()
	s.store.store.Set(s.id, values)
	return nil
}
