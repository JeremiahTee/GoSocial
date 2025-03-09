package main

import (
	"github.com/JeremiahTee/GoSocial/internal/env"
	"github.com/JeremiahTee/GoSocial/internal/store"
	"log"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
		store:  store.NewStorage(nil),
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
