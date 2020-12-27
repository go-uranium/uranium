package ushio

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/go-ushio/ushio/data"
	"github.com/go-ushio/ushio/utils/mdparse"
	"github.com/go-ushio/ushio/utils/render"
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
func Start(address string, conf *Config) error {
	if locked {
		return errors.New("one instance only")
	}
	locked = true
	defer func() {
		data.Quit()
		locked = false
	}()

	config = conf

	engine := render.New("./views/", ".html")

	err := data.Init("mysql", config.SQL)
	if err != nil {
		return err
	}

	app := fiber.New(fiber.Config{
		Views: engine,
		ErrorHandler: func(ctx *fiber.Ctx, e error) error {
			switch e.(type) {
			case *fiber.Error:
				return ctx.Render("_error", e, "_error")
			default:
				log.Println(e)
				return ctx.Render("_error",
					fiber.Map{"Code": 500, "Message": "An unexpected error occurred!"},
					"_error")
			}
		},
	})

	app.Static("/static/", config.Static)
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Redirect("/home", 307)
	})
	app.Get("/home", HomeHandler)
	app.Get("/u/:name", UserHandler)
	app.Get("/p/:post", PostHandler)
	app.Get("/login", LoginHandler)
	app.Post("/login", LoginPostHandler)
	app.Get("/sign_up", SignUpHandler)
	app.Get("/compose", ComposeHandler)
	app.Post("/compose", ComposePostHandler)

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

	app.Get("/*", func(c *fiber.Ctx) error {
		return fiber.NewError(404, "Not found!!1")
	})

	return app.Listen(address)
}
