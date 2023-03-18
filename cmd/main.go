package main

import (
	"log"
    "github.com/maxime-louis14/api-golang/database"
	"github.com/gofiber/fiber/v2"
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