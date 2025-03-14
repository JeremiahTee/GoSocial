package main

import (
	"github.com/JeremiahTee/GoSocial/internal/db"
	"github.com/JeremiahTee/GoSocial/internal/env"
	"github.com/JeremiahTee/GoSocial/internal/store"
	"log"
)

const version = "0.0.1"

//	@title			GoSocial API
//	@description	API for GoSocial, a social network in Go
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath					/v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description
func main() {
	cfg := config{
		addr:   env.GetString("ADDR", ":8080"),
		apiURL: env.GetString("EXTERNAL_URL", "localhost:8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}

	database, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	defer database.Close()
	log.Println("database connection pool established")

	storage := store.NewStorage(database)

	app := &application{
		config: cfg,
		store:  storage,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
