package session

import (
	"session/provider/memory"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

// Store represents a session store
type Store struct {
	provider Provider
}

// Config defines the config for middleware.
type Config struct {
	// Storage interface implemented by storage providers
	Provider Provider
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Provider: memory.New(), // WOOOHOO
}

// New creates a new middleware handler
func New(config ...Config) *Store {
	// Set default config
	cfg := ConfigDefault

	// Override config if provided
	if len(config) > 0 {
		cfg = config[0]

		if cfg.Provider == nil {
			cfg.Provider = ConfigDefault.Provider
		}
	}

	return &Store{
		provider: cfg.Provider,
	}
}

// Get ...
func (s *Store) Get(c *fiber.Ctx) *Session {
	newUser := false

	id := c.Cookies("session_id")
	if len(id) == 0 {
		id = utils.UUID()
		newUser = true
	}

	// base sess
	sess := &Session{
		store:   s,
		id:      id,
		newUser: newUser,
	}

	// create data or new data?
	if !newUser {
		s.provider.Get(id)
		data, err := sess.store.provider.Get(id)
		if err != nil {
			panic(err)
		}
		// sess.data = nil, data = dictpool from db
		if err := decode(sess.data, data); err != nil {
			panic(err)
		}
	}

	return sess
}
