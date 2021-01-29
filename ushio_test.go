package ushio_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-ushio/ushio"
	"github.com/go-ushio/ushio/data"
)

func TestUshio(t *testing.T) {
	fmt.Println(runtime.NumCPU())
	db, err := sql.Open("mysql", os.Getenv(`DATA_SOURCE_NAME`))
	if err != nil {
		panic(err)
	}

	u := ushio.New(db, data.SQLSentence(), &ushio.Config{
		SiteName: "Ushio",
		SendMail: func(dst string, token string) error {
			fmt.Println(dst, token)
			return nil
		},
	})

	engine := html.New("./views", ".html")
	engine.AddFunc("dateFormat", func(date time.Time) string {
		sub := time.Now().Sub(date)
		hours := sub.Hours()
		minutes := sub.Minutes()
		switch {
		case hours < 1:
			switch {
			case minutes < 1:
				return "recently"
			default:
				return fmt.Sprintf("%.0f minute(s) ago", minutes)
			}
		case hours < 24:
			return fmt.Sprintf("%.0f hour(s) ago", hours)
		default:
			return fmt.Sprintf("%.0f day(s) ago", hours/24)
		}
	})
	engine.AddFunc("numFormat", func(i int) string {
		switch {
		case i < 1000:
			return strconv.Itoa(i)
		case i < 1000000:
			return fmt.Sprintf("%.2fk", float64(i)/1000)
		default:
			return "1M+"
		}
	})
	engine.Reload(true)
	app := fiber.New(fiber.Config{
		Views: engine,
		ErrorHandler: func(ctx *fiber.Ctx, e error) error {
			switch e.(type) {
			case *fiber.Error:
				return ctx.Render("error", e, "error")
			default:
				log.Println(e)
				return ctx.Render("error",
					fiber.Map{"Code": 500, "Message": "An unexpected error occurred!"})
			}
		},
	})
	app.Static("/static/", "static/")

	app.Get("/test/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello")
	})

	u.Configure(app)

	app.Get("/*", func(c *fiber.Ctx) error {
		return fiber.NewError(404, "Not found!!1")
	})

	err = app.Listen(":8888")
	t.Error(err)
}
