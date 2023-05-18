package config

import (
	"fmt"
	"gofiber-gorm/src/pkg/constants"
	"gofiber-gorm/src/pkg/helpers"
	"strings"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors() cors.Config {
	// list allowed origins
	allowedOrigins := strings.Join(constants.AllowedOrigins(), ", ")

	logMessage := helpers.PrintLog("Cors", "Allowed Origins ( "+allowedOrigins+" )")
	fmt.Println(logMessage)

	config := cors.Config{
		AllowOrigins: allowedOrigins,
		// AllowMethods:  "GET, POST, HEAD, PUT, DELETE, PATCH",
		// AllowHeaders:  "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token",
		// ExposeHeaders: "Content-Length",
		// MaxAge:        86400,
	}

	return config
}
