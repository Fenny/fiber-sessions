package memory

import "sync"

// type Provider interface {
// 	Get(key string) ([]byte, error)
// 	Set(key string, val interface{}) error
// 	Delete(key string) error
// }

type Provider struct {
	db map[string][]byte
	mu sync.Mutex
}

// New returns a new memory provider configured
func New() *Provider {
	return &Provider{
		db: make(map[string][]byte),
	}
}

func (p *Provider) Get(key string) ([]byte, error) {
	// p.mu.Lock()
	// defer p.mu.Unlock()
	val, exist := p.db[key]
	if !exist {
		return nil, nil
	}
	return val, nil
}

func (p *Provider) Set(key string, val []byte) error {
	// p.mu.Lock()
	// defer p.mu.Unlock()
	p.db[key] = val
	return nil
}

func (p *Provider) Delete(key string) error {
	// p.mu.Lock()
	// defer p.mu.Unlock()
	delete(p.db, key)
	return nil
}
