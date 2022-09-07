package routes

import (
	"gofiber-gorm/src/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func RouteV1(app *fiber.App) {
	// group v1
	v1 := app.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")

		return c.Next()
	})

	// group route role
	role := v1.Group("/role")
	role.Get("/", controllers.GetAll)
	role.Post("/", controllers.CreateRole)
}
