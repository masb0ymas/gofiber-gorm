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
	authHandler := v1.Group("/auth")
	authHandler.Post("/sign-up", controllers.Register)
	authHandler.Post("/sign-in", controllers.Login)
	authHandler.Get("/verify-session", middlewares.AuthMiddleware(), controllers.VerifySession)

	// Role Route
	roleHandler := v1.Group("/role", middlewares.AuthMiddleware())
	roleHandler.Get("/", controllers.FindAllRole)
	roleHandler.Get("/:id", controllers.FindRoleById)
	roleHandler.Post("/", controllers.CreateRole)
	roleHandler.Put("/:id", controllers.UpdateRole)
	roleHandler.Put("/restore/:id", controllers.RestoreRoleById)
	roleHandler.Delete("/soft-delete/:id", controllers.SoftDeleteRoleById)
	roleHandler.Delete("/force-delete/:id", controllers.ForceDeleteRoleById)

	// Upload Route
	uploadHandler := v1.Group("/upload")
	uploadHandler.Get("/", controllers.FindAllUpload)
	uploadHandler.Get("/:id", controllers.FindUploadById)
	uploadHandler.Get("/presign/:keyFile", controllers.PresignedUploadURL)
	uploadHandler.Post("/", controllers.CreateUpload)
	uploadHandler.Put("/:id", controllers.UpdateUpload)
	uploadHandler.Put("/restore/:id", controllers.RestoreUploadById)
	uploadHandler.Delete("/soft-delete/:id", controllers.SoftDeleteUploadById)
	uploadHandler.Delete("/force-delete/:id", controllers.ForceDeleteUploadById)

	// Session Route
	sessionHandler := v1.Group("/session", middlewares.AuthMiddleware())
	sessionHandler.Get("/", controllers.FindAllSession)
	sessionHandler.Get("/:id", controllers.FindSessionById)
	sessionHandler.Post("/", controllers.CreateSession)
	sessionHandler.Put("/:id", controllers.UpdateSession)
	sessionHandler.Put("/restore/:id", controllers.RestoreSessionById)
	sessionHandler.Delete("/soft-delete/:id", controllers.SoftDeleteSessionById)
	sessionHandler.Delete("/force-delete/:id", controllers.ForceDeleteSessionById)

	// User Route
	userHandler := v1.Group("/user", middlewares.AuthMiddleware())
	userHandler.Get("/", controllers.FindAllUser)
	userHandler.Get("/:id", controllers.FindUserById)
	userHandler.Post("/", controllers.CreateUser)
	userHandler.Put("/:id", controllers.UpdateUser)
	userHandler.Put("/restore/:id", controllers.RestoreUserById)
	userHandler.Delete("/soft-delete/:id", controllers.SoftDeleteUserById)
	userHandler.Delete("/force-delete/:id", controllers.ForceDeleteUserById)
}
