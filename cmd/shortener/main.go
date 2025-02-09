package main

import (
	"github.com/iteplenky/url-shortener/internal/pkg/app"
	"log"
)

func main() {

	application, err := app.New()

	if err != nil {
		log.Fatal(err)
	}

	if err = application.Run(); err != nil {
		log.Fatal(err)
	}
}
