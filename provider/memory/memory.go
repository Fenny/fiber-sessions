package memory

import "sync"

// type Provider interface {
// 	Get(key string) ([]byte, error)
// 	Set(key string, val interface{}) error
// 	Delete(key string) error
// }

type Provider struct {
	db map[string]interface{}
	mu sync.Mutex
}

type item struct {
	data []byte
}

// New returns a new memory provider configured
func New() *Provider {
	return &Provider{
		db: new(map[string]interface{}),
	}
}

func (p *Provider) Get(key string) ([]byte, error) {
	mu.Lock()
	defer mu.Unlock()
	val, exist := p.db[key]
	if !exist {
		return nil, nil
	}
	item := val.(*item)
	return item.data, nil
}

func (p *Provider) Set(key string, val interface{}) error {
	mu.Lock()
	defer mu.Unlock()
	p.db[key] = val
	return nil
}

func (p *Provider) Delete(key string) error {
	mu.Lock()
	defer mu.Unlock()
	delete(p.db, key)
	return nil
}
