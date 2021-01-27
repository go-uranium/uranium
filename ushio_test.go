package ushio_test

import (
	"log"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func TestUshio(t *testing.T) {
	engine := html.New("./views", ".html")
	engine.Reload(true)
	engine.Debug(true)
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
	app.Static("/static/", "static/")

	app.Get("/*", func(c *fiber.Ctx) error {
		return fiber.NewError(404, "Not found!!1")
	})

	return app.Listen(address)
}
