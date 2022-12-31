package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/jonafll/go-fiber-insurance-api/config"
	"github.com/jonafll/go-fiber-insurance-api/routes"

	_ "github.com/jonafll/go-fiber-insurance-api/docs"
)

// @title Car Insurance API
// @version 1.0
// @description Car Insurance API with Go and Fiber.
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api/v1
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Failed to load .env file", err)
	}

	config := config.New()
	app := fiber.New(config)

	routes.Swagger(app)
	routes.Tariff(app)

	if err := app.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Fatalln("Failed to start server", err)
	}
}
