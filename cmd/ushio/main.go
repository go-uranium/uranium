package main

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/mailgun/mailgun-go/v3"

	"github.com/go-ushio/ushio"
	"github.com/go-ushio/ushio/data"
)

/*
!FOR TEST ONLY!
!STILL DEVELOPING!
*/

func main() {
	db, err := sql.Open("postgres", os.Getenv(`DATA_SOURCE_NAME`))
	if err != nil {
		panic(err)
	}

	u := ushio.New(db, data.SQLSentence(), &ushio.Config{
		SiteName: "Ushio",
		SendMail: func(dst string, token string) error {
			var yourDomain = "ushio.zincic.com"
			var privateAPIKey = os.Getenv("MAILGUN_API_SEC")
			mg := mailgun.NewMailgun(yourDomain, privateAPIKey)
			sender := "no-reply@ushio.zincic.com"
			subject := "Complete your sign-up by verifying your email"
			file, err := ioutil.ReadFile("./views/email.html")
			if err != nil {
				return err
			}
			tpl, err := template.New("email").Parse(string(file))
			if err != nil {
				return err
			}
			buf := &bytes.Buffer{}
			err = tpl.Execute(buf, fiber.Map{
				"URL": fmt.Sprintf("https://ushio.zincic.com/sign_up?token=%s&email=%s", token, url.QueryEscape(dst)),
			})
			if err != nil {
				return err
			}
			message := mg.NewMessage(sender, subject, buf.String(), dst)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			_, _, err = mg.Send(ctx, message)
			if err != nil {
				return err
			}
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
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Bye!")
		os.Exit(0)
	}
}
