package main

import (
	"gofiber-gorm/src/pkg/config"
	"gofiber-gorm/src/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"

	_ "gofiber-gorm/docs" // load API Docs files (Swagger)
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email your@mail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app := fiber.New()

	port := config.Env("APP_PORT")

	// default middleware
	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(helmet.New())
	app.Use(logger.New())
	app.Use(limiter.New(limiter.Config{Max: 100}))

	// static file
	app.Static("/", "./public")

	// Connect to the Database
	config.ConnectDB()

	// initial app route
	routes.InitialRoutes(app)

	// listening app
	log.Fatal(app.Listen(":" + port))
}
