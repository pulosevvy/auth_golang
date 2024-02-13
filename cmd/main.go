package main

import (
	"log"
	"test_app/config"
	"test_app/internal/app"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg)
}
