package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	inicializeFiberServer()
}

func inicializeFiberServer() {

	app := fiber.New()

	app.Use(cors.New(cors.ConfigDefault))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/api")
	})

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("api/messages", handleMessages)

	app.Listen(":3000")

}

func handleMessages(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	if data["name"] == "" {
		data["name"] = "World"
	}

	data["message"] = "Hello, " + data["name"]

	delete(data, "name")

	return c.JSON(data)
}
