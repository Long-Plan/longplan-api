package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello world!")
	})
	app.Get("/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.Status(200).SendString("Hello " + name)
	})
	app.Listen(":8000")
}
