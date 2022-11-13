package config

import "github.com/gofiber/fiber/v2/middleware/cors"

func Cors() cors.Config {
	config := cors.Config{
		AllowOrigins: "http://localhost:3000, http://localhost:3333",
		// AllowMethods:  "GET, POST, HEAD, PUT, DELETE, PATCH",
		// AllowHeaders:  "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token",
		// ExposeHeaders: "Content-Length",
		// MaxAge:        86400,
	}

	return config
}
