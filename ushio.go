package ushio

import (
	"sync"

	"github.com/gofiber/fiber/v2"

	"github.com/go-ushio/ushio/cache"
	"github.com/go-ushio/ushio/data"
	"github.com/go-ushio/ushio/utils/mdparse"
)

type Config struct {
	SiteName string
	SendMail func(dst string, token string) error
}

// call ushio.Lock.Lock() before exiting
type Ushio struct {
	Data   data.Provider
	Cache  cache.Cacher
	Config *Config
	Lock   *sync.RWMutex
}

func New(provider data.Provider, config *Config) *Ushio {
	return &Ushio{
		Data:   provider,
		Cache:  cache.New(provider),
		Config: config,
		Lock:   &sync.RWMutex{},
	}
}

func (ushio *Ushio) Configure(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Redirect("/home", 307)
	})
	app.Get("/home", ushio.HomeHandler)
	app.Get("/u/:name", ushio.UserHandler)
	app.Get("/p/:post", ushio.PostHandler)
	app.Get("/login", ushio.LoginHandler)
	app.Post("/login", ushio.LoginPostHandler)
	app.Get("/sign_up", ushio.SignUpHandler)
	app.Post("/sign_up", ushio.SignUpPostHandler)
	app.Get("/compose", ushio.ComposeHandler)
	app.Post("/compose", ushio.ComposePostHandler)
	app.Get("/logout", func(ctx *fiber.Ctx) error {
		ctx.ClearCookie("token")
		return ctx.Redirect("/", 307)
	})

	app.Post("/md_parse", func(ctx *fiber.Ctx) error {
		html, e := mdparse.Parse(string(ctx.Body()))
		if e != nil {
			return e
		}
		return ctx.SendString(string(*html))
	})
}
