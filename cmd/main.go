package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/maxime-louis14/api-golang/database"
	"github.com/maxime-louis14/api-golang/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awsome API")
}

func setupRoutes(app *fiber.App) {
	//
	app.Get("/api", welcome)
	//User endpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/user/:id", routes.DeleteUser)
	// Product endpoints
	app.Post("api/recettes", routes.CreateRecette)
	app.Get("/api/recettes", routes.GetRecettes)
	app.Get("/api/recettes/:id", routes. GetRecette)
	app.Put("/api/recettes/:id", routes.UpdateRecette)
	app.Delete("/api/recette/:id", routes.DeleteRecette)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
