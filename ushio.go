package ushio

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	"github.com/go-ushio/ushio/data"
)

var (
	config *Config
	locked bool
)

// Start starts an instance of ushio.
// You can pass an optional *tls.Config to enable TLS.
//
// u.Start(8080)
// u.Start("8080")
// u.Start(":8080")
// u.Start("127.0.0.1:8080")
func Start(address string, config *Config) error {
	if locked {
		return errors.New("one instance only")
	}
	locked = true
	defer func() {
		data.Quit()
		locked = false
	}()

	engine := html.New("./views", ".html")

	err := data.Init("mysql", config.SQL)
	if err != nil {
		return err
	}

	app := fiber.New()
	app.Static("/static/", config.Static)
	app.Get("/u/:name", UserHandler)
	return app.Listen(address)
}
