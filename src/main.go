package main

import (
	"log"

	"github.com/pedro-git-projects/nilptr/src/app"
)

func main() {
	if err := app.New().Start(); err != nil {
		log.Fatal(err)
	}
}
