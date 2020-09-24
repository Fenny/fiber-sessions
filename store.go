package session

import (
	"github.com/hi019/fiber-sessions/provider/memory"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

// Storage interface implemented by providers
type Storage interface {
	// Set session value, call save function to take effect
	Set(key string, value interface{})
	// Get session value
	Get(key string) (interface{}, bool)
	// Delete session value, call save function to take effect
	Delete(key string)
}

// Store represents a session store
type Store struct {
	store Storage
}

// Config defines the config for middleware.
type Config struct {
	// Storage interface implemented by storage providers
	Storage Storage
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Storage: memory.New(), // WOOOHOO
}

// New creates a new middleware handler
func New(config ...Config) *Store {
	// Set default config
	cfg := ConfigDefault

	// Override config if provided
	if len(config) > 0 {
		cfg = config[0]

		if cfg.Storage == nil {
			cfg.Storage = ConfigDefault.Storage
		}
	}

	return &Store{cfg.Storage}
}

// Get ...
func (s *Store) Get(c *fiber.Ctx) *Session {
	var new bool

	// Get ID from cookie
	id := c.Cookies("session_id")

	// If no ID exist, create new one
	if len(id) == 0 {
		id = utils.UUID()
		new = true
	}

	// Create session object
	sess := &Session{
		ctx:   c,
		store: s,
		new:   new,
		id:    id,
	}

	// Fetch existing data
	if !new {
		raw, found := s.store.Get(id)
		// Set data
		if found {
			sess.values = raw.(map[string]interface{})
			return sess
		}
	}

	// Create new storage TODO: Pool
	sess.values = make(map[string]interface{})

	// Return session object
	return sess
}
