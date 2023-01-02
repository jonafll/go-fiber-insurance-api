package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonafll/go-fiber-insurance-api/handler"
)

func Vehicle(a *fiber.App) {
	route := a.Group("/api/v1")
	route.Get("/vehicles", handler.GetVehicles)
	route.Get("/vehicles/:id<int>;<max(999)>", handler.GetVehicleWithId)
}
