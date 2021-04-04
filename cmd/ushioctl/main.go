package main

import (
	"fmt"
	"os"
)

var helpMsg = ``

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}

	//dsn := os.Getenv(`postgres://ushio:9rShFGyzD4oHxA4qmEzmgfmag79yxJq2@127.0.0.1/postgres?sslmode=disable`)

	switch os.Args[1] {
	case "dump":

	default:
		help()
		return
	}
}

func help() {
	fmt.Println(helpMsg)
}
