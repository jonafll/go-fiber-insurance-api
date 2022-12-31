package handler

import "github.com/gofiber/fiber/v2"

// @Description Start tariffing.
// @Summary Start tariffing
// @Tags Tariffing
// @Accept json
// @Produce json
// @Success 200 {string} status "ok"
// @Failure 400 {string} status "bad request"
// @Failure 500 {string} status "internal server error"
// @Router /tariff [post]
func Tariff(c *fiber.Ctx) error {
	return c.JSON("{}")
}
