package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonafll/go-fiber-insurance-api/handler"
)

func Tariff(a *fiber.App) {
	route := a.Group("/api/v1")
	route.Post("/tariff", handler.Tariff)
}
