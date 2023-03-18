package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/maxime-louis14/api-golang/database"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awsome API")
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":3000"))
}
