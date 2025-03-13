package main

import (
	"github.com/JeremiahTee/GoSocial/internal/db"
	"github.com/JeremiahTee/GoSocial/internal/env"
	"github.com/JeremiahTee/GoSocial/internal/store"
	"log"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	store := store.NewStorage(conn)
	db.Seed(store)
}
