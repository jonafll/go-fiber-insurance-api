package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/jonafll/go-fiber-insurance-api/config"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Failed to load .env file", err)
	}

	config := config.New()
	app := fiber.New(config)

	if err := app.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Fatalln("Failed to start server", err)
	}
}
