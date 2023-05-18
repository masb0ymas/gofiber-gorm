package controllers

import (
	"gofiber-gorm/src/app/service"
	"gofiber-gorm/src/database/schema"
	"gofiber-gorm/src/pkg/config"
	"gofiber-gorm/src/pkg/helpers"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Find All Session
func FindAllSession(c *fiber.Ctx) error {
	db := config.GetDB()

	sessionService := service.NewSessionService(db)
	data, total, err := sessionService.FindAll(c)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "error to get sessions",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been received",
		"data":    data,
		"total":   total,
	})
}

// Find Session By Id
func FindSessionById(c *fiber.Ctx) error {
	db := config.GetDB()

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to get session",
		})
	}

	sessionService := service.NewSessionService(db)
	data, err := sessionService.FindById(id)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"code":    http.StatusNotFound,
			"message": "data not found or has been deleted",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been received",
		"data":    data,
	})
}

// Create Session
func CreateSession(c *fiber.Ctx) error {
	db := config.GetDB()

	sessionSchema := new(schema.SessionSchema)

	if err := helpers.ParseFormDataAndValidate(c, sessionSchema); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.NewError(http.StatusBadRequest))
	}

	// session service
	sessionService := service.NewSessionService(db)
	data, err := sessionService.Create(*sessionSchema)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to create session",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been added",
		"data":    data,
	})
}

// Update Session
func UpdateSession(c *fiber.Ctx) error {
	db := config.GetDB()

	sessionSchema := new(schema.SessionSchema)

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to get session",
		})
	}

	if err := helpers.ParseFormDataAndValidate(c, sessionSchema); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.NewError(http.StatusBadRequest))
	}

	// session service
	sessionService := service.NewSessionService(db)
	data, err := sessionService.Update(id, *sessionSchema)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to update session",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been updated",
		"data":    data,
	})
}

// Restore
func RestoreSessionById(c *fiber.Ctx) error {
	db := config.GetDB()

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to get session",
		})
	}

	sessionService := service.NewSessionService(db)
	err = sessionService.Restore(id)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"code":    http.StatusNotFound,
			"message": "data not found or has been deleted",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been updated",
	})
}

// Soft Delete
func SoftDeleteSessionById(c *fiber.Ctx) error {
	db := config.GetDB()

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to get session",
		})
	}

	sessionService := service.NewSessionService(db)
	err = sessionService.SoftDelete(id)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"code":    http.StatusNotFound,
			"message": "data not found or has been deleted",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been deleted",
	})
}

// Force Delete
func ForceDeleteSessionById(c *fiber.Ctx) error {
	db := config.GetDB()

	id, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"code":    http.StatusBadRequest,
			"message": "failed to get session",
		})
	}

	sessionService := service.NewSessionService(db)
	err = sessionService.ForceDelete(id)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"code":    http.StatusNotFound,
			"message": "data not found or has been deleted",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"code":    http.StatusOK,
		"message": "data has been deleted",
	})
}
