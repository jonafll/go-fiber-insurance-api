package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/jonafll/go-fiber-insurance-api/config"
	"github.com/jonafll/go-fiber-insurance-api/middleware"
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

	middleware.Fiber(app)

	routes.Swagger(app)
	routes.Tariff(app)

	go func() {
		if err := app.Listen(os.Getenv("SERVER_URL")); err != nil {
			log.Fatalln("Failed to start server", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c
	log.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	log.Println("Running cleanup tasks...")
	// Your cleanup tasks go here
	log.Println("Fiber was successful shutdown.")
}
