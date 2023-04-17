package main

import (
	"log"
	"os"

	"github.com/alvinmdj/mygram-api/database"
	"github.com/alvinmdj/mygram-api/helpers"
	"github.com/alvinmdj/mygram-api/routers"
	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func main() {
	database.StartDB()
	helpers.InitCloudinary()
	r := routers.StartApp()
	r.Run()
}
