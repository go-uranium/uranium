package route

import (
	"github.com/gofiber/fiber"

	"github.com/go-ushio/ushio/view"
)

func Set(r fiber.Router) {
	r.Get("/",view.HandleIndex)

	r.Get("/api/user/:uid",view.HandleUserAPI)
}