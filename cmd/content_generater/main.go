package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/go-ushio/ushio/data/postgres"
	"github.com/go-ushio/ushio/model/category"
	"github.com/go-ushio/ushio/model/post"
	"github.com/go-ushio/ushio/model/user"
)

type Faker struct {
	*gofakeit.Faker
}

func main() {
	faker := Faker{gofakeit.New(time.Now().UnixNano())}
	posts := faker.GeneratePosts(5000)

	pg, err := postgres.New(os.Getenv(`DATA_SOURCE_NAME`))
	if err != nil {
		log.Fatalln(err)
		return
	}

	for i := range posts {
		_, err := pg.InsertPost(posts[i])
		if err != nil {
			log.Fatalln(err)
			return
		}
		fmt.Println(i)
	}

	//for i := 7; i < 600; i++ {
	//	_, err := pg.InsertUser(faker.GenerateUser())
	//	if err != nil {
	//		log.Fatalln(err)
	//		return
	//	}
	//	fmt.Println(i)
	//}

}

//Title
//Creator
//CreatedA
//LastMod
//Replies
//Views
//Activity
//Hidden
//// uid l
//VotePos
//VoteNeg
//Limit
//Category

func (faker *Faker) GeneratePosts(size int) []*post.Post {
	posts := make([]*post.Post, size)
	for i := range posts {
		fmt.Println(i)
		posts[i] = &post.Post{
			Info: &post.Info{
				Title: faker.Sentence(faker.Number(2, 25)),
				Creator: user.Simple{
					UID: int64(faker.Number(0, 100)),
				},
				CreatedAt: faker.Date(),
				LastMod:   faker.Date(),
				Replies:   int64(faker.Number(0, 500)),
				Views:     int64(faker.Number(0, 5000)),
				Activity:  faker.Date(),
				Hidden:    faker.Bool(),
				Category: category.Category{
					TID: int64(faker.Number(0, 5)),
				},
			},
			Content: template.HTML(faker.Paragraph(3, 3, 100, "<br/>")),
		}
	}
	return posts
}

func (faker *Faker) GenerateUserSimple() *user.Simple {
	return &user.Simple{
		Name:     faker.Name(),
		Username: faker.Username(),
		Avatar:   faker.LetterN(32),
	}
}

func (faker *Faker) GenerateUser() *user.User {
	r := make([]byte, 10)
	_, _ = rand.Read(r)
	rm := md5.Sum(r)
	return &user.User{
		Name:      faker.Name(),
		Username:  faker.Username(),
		Email:     faker.Email(),
		Avatar:    hex.EncodeToString(rm[:]),
		Bio:       faker.Quote(),
		CreatedAt: faker.Date(),
		Artifact:  int64(faker.Number(0, 1000)),
	}
}

//func (faker *Faker) GenerateUserAuth() []*user.Auth {
//
//}
//
//
//func (faker *Faker) GenerateUsers(size int) []*user.User {
//
//}
//
//func (faker *Faker) GenerateUserAuths(size int) []*user.Auth {
//
//}
