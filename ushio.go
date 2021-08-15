package uranium

import (
	"sync"

	"github.com/gofiber/fiber/v2"

	"github.com/go-uranium/uranium/utils/sendmail"

	"github.com/go-uranium/uranium/cache"
	"github.com/go-uranium/uranium/data"
	"github.com/go-uranium/uranium/utils/mdparse"
)

type Config struct {
	SiteName string
	Sender   sendmail.Sender
	PageSize int64
}

// call uranium.Lock.Lock() before exiting
type Ushio struct {
	Data   data.Provider
	Cache  cache.Cacher
	Config *Config
	Lock   *sync.RWMutex
}

func New(provider data.Provider, config *Config) (*Ushio, error) {
	if config.PageSize < 1 {
		config.PageSize = 35
	}
	cc := cache.New(provider, config.PageSize)
	err := cc.Init()
	if err != nil {
		return &Ushio{}, err
	}
	return &Ushio{
		Data:   provider,
		Cache:  cc,
		Config: config,
		Lock:   &sync.RWMutex{},
	}, nil
}

func (ushio *Ushio) Configure(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Redirect("/home", 307)
	})
	app.Get("/home", ushio.HandleHome)
	app.Get("/u/:name", ushio.HandleUser)
	app.Get("/u/:name/posts", func(ctx *fiber.Ctx) error {
		return ctx.Redirect("/u/"+ctx.Params("name"), 302)
	})
	app.Get("/u/:name/comments", ushio.HandleUserComments)
	app.Get("/p/:post", ushio.HandlePost)
	app.Get("/c/:tname", ushio.HandleCategory)
	app.Get("/login", ushio.HandleLogin)
	app.Post("/login", ushio.HandlePOSTLogin)
	app.Get("/sign_up", ushio.HandleSignUp)
	app.Post("/sign_up", ushio.HandlePOSTSignUp)
	app.Get("/compose", ushio.HandleCompose)
	app.Post("/compose", ushio.HandlePOSTCompose)
	app.Post("/vote/post", ushio.HandlePOSTVotePost)
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
