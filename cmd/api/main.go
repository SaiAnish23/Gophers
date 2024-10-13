package main

import (
	"log"

	"github.com/SaiAnish23/Gophers/internal/env"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config{
		addr: env.GetString("ADDR", ":4000"),
	}

	app := &application{
		config: cfg,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
