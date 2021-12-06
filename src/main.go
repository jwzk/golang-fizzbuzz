package main

import (
	"fizzbuzz/core/app"
	"fizzbuzz/core/web"
	"log"

	"github.com/joho/godotenv"
)

func webAction() {
	web.InitServer(app.Router)
}

func setup() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	setup()
	webAction()
}
