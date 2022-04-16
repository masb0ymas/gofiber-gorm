package routes

import (
	"gofiber-gorm/src/controllers"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func RouteV1(app *fiber.App) {
	// group v1
	v1 := app.Group("/v1")

	// group route role
	role := v1.Group("/role")
	role.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.NewError(http.StatusOK))
	})
	role.Post("/", controllers.CreateRole)
}
