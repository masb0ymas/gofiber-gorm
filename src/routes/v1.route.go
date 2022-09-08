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

	// Role Route
	roleRoute := v1.Group("/role")
	roleRoute.Get("/", controllers.FindAllRole)
	roleRoute.Get("/:id", controllers.FindRoleById)
	roleRoute.Post("/", controllers.CreateRole)
	roleRoute.Put("/:id", controllers.UpdateRole)
	roleRoute.Put("/restore/:id", controllers.RestoreRoleById)
	roleRoute.Delete("/soft-delete/:id", controllers.SoftDeleteRoleById)
	roleRoute.Delete("/force-delete/:id", controllers.ForceDeleteRoleById)
}
