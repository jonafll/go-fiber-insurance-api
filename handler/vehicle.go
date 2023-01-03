package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonafll/go-fiber-insurance-api/model"
	"github.com/jonafll/go-fiber-insurance-api/utils"
)

// @Description Search vehicle.
// @Summary Search vehicle
// @Tags Vehicle
// @Produce json
// @Param manufacturer query string true "Manufacturer"
// @Param model query string true "Model"
// @Param limit query int false "limit"
// @Success 200 {array} model.Vehicle
// @Failure 400 {object} utils.httpError
// @Failure 500 {object} utils.httpError
// @Router /vehicles [get]
func GetVehicles(c *fiber.Ctx) error {
	utils.SimulateRestCall()

	// Dummy for now
	vehicles := &[]model.Vehicle{
		{
			Id:           1,
			Manufacturer: "VW",
			Model:        "ID.3",
			Performance:  200,
			Vin:          "D385DHDJ239484ASQ",
		},
		{
			Id:           2,
			Manufacturer: "Audi",
			Model:        "RS6",
			Performance:  430,
			Vin:          "D474746GSJSTQ7B99",
		},
	}

	return c.JSON(vehicles)
}

// @Description Search vehicle with id.
// @Summary Search vehicle with id
// @Tags Vehicle
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} model.Vehicle
// @Failure 400 {object} utils.httpError
// @Failure 500 {object} utils.httpError
// @Router /vehicles/{id} [get]
func GetVehicleWithId(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "id must be of type integer")
	}

	utils.SimulateRestCall()

	// Dummy for now
	vehicle := &model.Vehicle{
		Id:           id,
		Manufacturer: "VW",
		Model:        "ID.3",
		Performance:  200,
		Vin:          "D385DHDJ239484ASQ",
	}

	return c.JSON(vehicle)
}
