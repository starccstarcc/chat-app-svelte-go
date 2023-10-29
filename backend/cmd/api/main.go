package main

import (
	"github.com/aletomasella/svelte-go-chat/pkg/store"
	"github.com/aletomasella/svelte-go-chat/pkg/store/postgres"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db, err := store.ConnectToDB("postgres", postgres.PostgresDNS)

	if err != nil {
		panic(err)
	}

	db.Ping()

	print("Connected to DB")
	// store := postgres.NewStore(db)

	inicializeServer()

}

func inicializeServer() {

	app := fiber.New()

	app.Use(cors.New(cors.ConfigDefault))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/api")
	})

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")

}
