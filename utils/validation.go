package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

var _ = validate.RegisterValidation("futuredate", futureDate)

func Validate(payload interface{}) *fiber.Error {
	err := validate.Struct(payload)
	if err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(
				errors,
				fmt.Sprintf("`%v` with value `%v` doesn't satisfy the `%v` constraint", err.Field(), err.Value(), err.Tag()),
			)
		}

		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: strings.Join(errors, ","),
		}
	}

	return nil
}

func futureDate(fl validator.FieldLevel) bool {
	t, err := time.Parse("2006-01-02", fl.Field().String())
	return err == nil && t.After(time.Now())
}
