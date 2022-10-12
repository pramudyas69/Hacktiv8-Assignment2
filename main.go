package main

import (
	"Assignment2/router"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error load .env file", err)
	}
	PORT := os.Getenv("PORT")

	App := router.StartApp()

	err = App.Run((PORT))
	if err != nil {
		log.Fatal("Error start the server", err)
	}
}
