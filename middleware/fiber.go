package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func Fiber(a *fiber.App) {
	a.Use(
		cors.New(),
		requestid.New(),
		logger.New(logger.Config{
			Format: "${locals:requestid} ${status} - ${method} ${path}\n",
		}),
	)
}
