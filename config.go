package ushio

import (
	"crypto/tls"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	SiteName        string
	TLS             *tls.Config
	Static          string
	StaticSettings  *fiber.Static
	SQL             string
	TemplatesReload bool
	TemplatesDebug  bool
}

var DefaultConfig = &Config{
	SiteName:        "Ushio",
	TLS:             nil,
	Static:          "static",
	StaticSettings:  nil,
	SQL:             os.Getenv("SQL"),
	TemplatesReload: false,
	TemplatesDebug:  false,
}
