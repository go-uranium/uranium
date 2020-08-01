package ushio

import (
	"log"
	"os"

	"github.com/gofiber/fiber"

	"github.com/go-ushio/ushio/route"
	"github.com/go-ushio/ushio/utils/render"
)

type Ushio struct {
	App *fiber.App
	Layout string
	Static string
	Logger *log.Logger
}

func NewUshio(app *fiber.App) *Ushio {
	return &Ushio{
		App:app,
		Layout:"layout/",
		Static:"static/",
		Logger:log.New(os.Stdout,"",log.Flags()),
	}
}

func (u *Ushio)Start() error {
	u.Logger.Println("Ushio is Starting...")

	// set renderer
	u.App.Settings.Views = render.New("layout","layout/partials",".html")

	// set route
	route.Set(u.App)

	u.App.Static("static","static")

	u.Logger.Println("Ushio is Running!")
	return u.App.Listen(":8044")
}