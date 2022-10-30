package routes

import (
	"gofiber-gorm/src/app/controllers"
	"gofiber-gorm/src/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RouteV1(app *fiber.App) {
	// group v1
	v1 := app.Group("/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")

		return c.Next()
	})

	// Auth Route
	authRoute := v1.Group("/auth")
	authRoute.Post("/sign-up", controllers.Register)
	authRoute.Post("/sign-in", controllers.Login)

	// Role Route
	roleRoute := v1.Group("/role", middlewares.AuthMiddleware())
	roleRoute.Get("/", controllers.FindAllRole)
	roleRoute.Get("/:id", controllers.FindRoleById)
	roleRoute.Post("/", controllers.CreateRole)
	roleRoute.Put("/:id", controllers.UpdateRole)
	roleRoute.Put("/restore/:id", controllers.RestoreRoleById)
	roleRoute.Delete("/soft-delete/:id", controllers.SoftDeleteRoleById)
	roleRoute.Delete("/force-delete/:id", controllers.ForceDeleteRoleById)

	// User Route
	userRoute := v1.Group("/user")
	userRoute.Get("/", middlewares.AuthMiddleware(), controllers.FindAllUser)
}
