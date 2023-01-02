package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonafll/go-fiber-insurance-api/model"
	"github.com/jonafll/go-fiber-insurance-api/utils"
)

// @Description Start tariffing.
// @Summary Start tariffing
// @Tags Tariffing
// @Accept json
// @Produce json
// @Param tariff body model.Tariff true "Tariff"
// @Success 200 {object} model.TariffAmount
// @Failure 400 {object} utils.httpError
// @Failure 500 {object} utils.httpError
// @Router /tariff [post]
func Tariff(c *fiber.Ctx) error {
	tariff := &model.Tariff{}

	if err := utils.ParseBodyAndValidate(c, tariff); err != nil {
		return err
	}

	utils.SimulateRestCall()

	tariffAmount := &model.TariffAmount{
		Total:         122.43,
		Partial:       32.65,
		Comprehensive: 76.43,
		Liability:     32.44,
		Currency:      "EUR",
	}

	return c.JSON(tariffAmount)
}
