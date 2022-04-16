package helpers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// validate input
func Validate(payload interface{}) []*fiber.Error {
	err := validate.Struct(payload)

	if err != nil {
		var errorList []*fiber.Error

		for _, err := range err.(validator.ValidationErrors) {
			errorList = append(errorList, &fiber.Error{
				Code:    fiber.StatusBadRequest,
				Message: fmt.Sprintf("%v must be valid", err.StructField()),
			})
		}

		return errorList
	}

	return nil
}

// parse form data body
func ParseFormData(c *fiber.Ctx, body interface{}) []*fiber.Error {
	if err := c.BodyParser(body); err != nil {
		var errorList []*fiber.Error
		errorList = append(errorList, fiber.ErrUnprocessableEntity)

		return errorList
	}

	return nil
}

// parse form data body and validate form
func ParseFormDataAndValidate(c *fiber.Ctx, body interface{}) []*fiber.Error {
	if err := ParseFormData(c, body); err != nil {
		return err
	}

	return Validate(body)
}
