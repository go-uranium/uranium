package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"text/template"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"

	"github.com/go-uranium/uranium/utils/sendmail"

	"github.com/go-uranium/uranium"
	"github.com/go-uranium/uranium/storage/postgres"
)

/*
!FOR TEST ONLY!
!STILL DEVELOPING!
*/

func main() {
	pg, err := postgres.New(os.Getenv(`DATA_SOURCE_NAME`))
	if err != nil {
		panic(err)
	}

	textRender, err := template.New("email.txt.views").ParseFiles("views/email.txt.views")
	if err != nil {
		log.Fatalln(err)
		return
	}

	u, err := uranium.New(pg, &uranium.Config{
		SiteName: "Ushio",
		Sender: &sendmail.SMTPClient{
			From:     "no-reply@uranium.zincic.com",
			Password: os.Getenv("SMTP_PASSWORD"),
			Host:     "smtp.mailgun.org",
			Port:     "587",
			Subject:  "Verify your email address.",
			Text:     textRender,
		},
	})

	if err != nil {
		log.Fatalln(err)
		return
	}

	engine := html.New("./views", ".views")
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
	engine.AddFunc("numFormat", func(i int64) string {
		switch {
		case i < 1000:
			return strconv.Itoa(int(i))
		case i < 1000000:
			return fmt.Sprintf("%.2fk", float64(i)/1000)
		default:
			return "1M+"
		}
	})
	// set to false to get better performance
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
		//Prefork:true,
	})
	app.Static("/", "static/")

	app.Get("/test/", func(ctx *fiber.Ctx) error {
		u.Lock.RLock()
		defer u.Lock.RUnlock()
		time.Sleep(10 * time.Second)
		return ctx.SendString("hello\n")
	})

	u.Configure(app)

	app.Get("/*", func(c *fiber.Ctx) error {
		return fiber.NewError(404, "Not found!!1")
	})

	go func() {
		err = app.Listen(":8888")
		log.Fatal(err)
	}()
	sgn := make(chan os.Signal, 1)
	signal.Notify(sgn, os.Interrupt, os.Kill)
	for range sgn {
		fmt.Println("Exiting... Please wait...")
		u.Lock.Lock()
		err := pg.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Bye!")
		os.Exit(0)
	}
}
