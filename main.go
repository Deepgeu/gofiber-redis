package main

import (
	"gofiber-redis/db"
	"gofiber-redis/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db.ConnectRedis()

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🚀 Redis Server Running")
	})

	app.Get("/seed", handlers.SeedRecords)
	app.Get("/record/:key", handlers.GetRecord)
	app.Post("/record", handlers.CreateRecord)
	app.Put("/record/:key", handlers.UpdateRecord)
	app.Delete("/record/:key", handlers.DeleteRecord)

	app.Listen(":3000")
}
