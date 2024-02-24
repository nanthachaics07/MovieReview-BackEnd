package router

import "github.com/gofiber/fiber"

func RouterBrowser(app *fiber.App) {
	// TODO: implement
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
