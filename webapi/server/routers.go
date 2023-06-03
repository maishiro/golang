package server

import (
	"github.com/gofiber/fiber/v2"

	ctrl "webapi/controller"
)

func NewRouter() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/v2", Index)
	app.Post("/v2/user", ctrl.CreateUser)
	app.Get("/v2/user/:username", ctrl.GetUserByName)

	return app
}

func Index(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
