package main

import (
	"fmt"
	"os"
	"time"

	"github.com/andikaraditya/personal-library/internal/user"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
)

func main() {
	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:        "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${queryParams} | ${error}\n",
		TimeFormat:    "15:04:05",
		TimeZone:      "Local",
		TimeInterval:  500 * time.Millisecond,
		Output:        os.Stdout,
		DisableColors: false,
	}))

	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/status", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": true,
		})
	})

	app.Post("/register", user.CreateUser)
	app.Post("login", user.Login)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(cfg.JWTSecret)},
	}))

	log.Fatal().Err(app.Listen(fmt.Sprintf(":%d", cfg.Port))).Msg("Server has crashed")
}
