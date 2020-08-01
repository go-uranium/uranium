package bind

import (
	"github.com/gofiber/fiber"

	"github.com/go-ushio/ushio/config"
	"github.com/go-ushio/ushio/controllor"
)

type Index struct {
	Site *config.Site
	Session *controllor.Session
}

func IndexData(c *fiber.Ctx) (*Index,error) {

	return &Index{
		Site:config.SiteConf,
		Session:&controllor.Session{
		LoggedIn:true,
		User:&controllor.User{
			Username:"admin",
		},
		},
	},nil
}