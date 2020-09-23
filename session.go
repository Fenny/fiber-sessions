package session

import "github.com/savsgio/dictpool"

// Session represents a single session
type Session struct {
	store   *Store
	data    *dictpool.Dict
	id      string
	newUser bool
}

// ID ...
func (s *Session) ID() string {
	return s.id
}

// Set ...
func (s *Session) Set(key string, val interface{}) {
	s.data.Set(key, val)
}

// Get ...
func (s *Session) Get(key string) interface{} {
	return s.data.Get(key)
}

// Delete ...
func (s *Session) Delete(key string) {
	s.data.Del(key)
}

// Destroy ...
func (s *Session) Destroy() {
	s.data.Reset()
}

// Save ...
func (s *Session) Save() error {
	raw, err := encode(*s.data)
	if err != nil {
		return err
	}
	return s.store.provider.Set(s.id, raw)
}

// IsNew ...
func (s *Session) IsNew() bool {
	return s.newUser
}

// Thanks to fasthttp/session for the below methods :D

// convert dictpool.Dict to []byte
func encode(src dictpool.Dict) ([]byte, error) {
	if len(src.D) == 0 {
		return nil, nil
	}
	dst, err := src.MarshalMsg(nil)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

// convert []byte to dictpool.Dict
func decode(dst *dictpool.Dict, src []byte) error {
	dst.Reset()
	if len(src) == 0 {
		return nil
	}
	_, err := dst.UnmarshalMsg(src)
	return err
}
