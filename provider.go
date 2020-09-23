package session

// Provider interface implemented by providers
type Provider interface {
	Get(key string) (interface{}, error)
	Set(key string, val interface{}) error
	Delete(key string) error
}

// func demo() {
// 	// mysql setup stuff

// 	store := session.New()

// 	app.Get("/", func(c *fiber.Ctx) error {
// 		sess := store.Get(c)
// 		if sess.Get("secret_key") != "lol123" {
// 			return c.SendStatus(403)
// 		}

// 		sess.ID()
// 		sess.Set("name", []byte("fenny"))
// 		sess.Set("name", []byte("fenny"))
// 		sess.Set("name", []byte("fenny"))
// 		sess.Set("name", []byte("fenny"))
// 		sess.Set("name", []byte("fenny"))
// 		sess.Set("name", []byte("fenny"))
// 		sess.Get("name")
// 		sess.Delete("name")
// 		sess.Destroy()
//		sess.Save()

// 		return c.SendString("Hello, World 👋")

// 	})
// }

// https://github.com/gofiber/session/v2
// https://github.com/gofiber/session/v2/provider/mysql

// https://github.com/gofiber/fiber/v2/middleware/session
// https://github.com/gofiber/storage/mysql
// https://github.com/gofiber/storage/redis
// https://github.com/gofiber/storage/postgres

// https://github.com/expressjs/session

// sessions := session.New()

// app.Get("/", func(c *fiber.Ctx) error {
// 	store := sessions.Get(c) // get new session
//  if store.ID() != "" {
// 		// some operations
// 		// query ID to DB
// 		//
// 		store.Set()
// }

// 	store := sessions.Set(c) // set new session
// 	store.Save()

// 	store.ID()               // returns session id
// 	store.Destroy()          // delete storage + cookie
// 	store.Get("john")        // get from storage
// 	store.Regenerate()       // generate new session id
// 	store.Delete("john")     // delete from storage
// 	store.Set("john", "doe") // save to storage

// 	return nil
// })

// Session.Get
// Session.Delete
// Session.Set

// session.regenerate
// session.destroy
// session.reload
// session.save
// session.touch
// session.sessionID

// store.all
// store.destroy
// store.clear
// store.length
// store.get
// store.touch
