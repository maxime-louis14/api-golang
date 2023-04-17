package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/maxime-louis14/api-golang/database"
	"github.com/maxime-louis14/api-golang/models"
)

type Recette struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Descriptions string `json:"descriptions"`
	Ingredients  string `json:"ingredients"`
	Photos       string `json:"photos"`
	Directions   string `json:"directions"`
	Page         string `json:"line"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseRecette(recetteModel models.Recette) Recette {
	return Recette{ID: recetteModel.ID, Name: recetteModel.Name, Descriptions: recetteModel.Descriptions, Ingredients: recetteModel.Ingredients, Photos: recetteModel.Photos, Directions: recetteModel.Directions, Page: recetteModel.Page, SerialNumber: recetteModel.SerialNumber}
}

func CreateRecette(c *fiber.Ctx) error {
	var recette models.Recette
	if err := c.BodyParser(&recette); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&recette)
	responseRecette := CreateResponseRecette(recette)
	return c.Status(200).JSON(responseRecette)
}

func GetRecettes(c *fiber.Ctx) error {
	recettes := []models.Recette{}
	database.Database.Db.Find(&recettes)
	responseRecettes := []Recette{}
	for _, recette := range recettes {
		responseRecette := CreateResponseRecette(recette)
		responseRecettes = append(responseRecettes, responseRecette)
	}
	return c.Status(200).JSON(responseRecettes)
}

func findRecette(id int, recette *models.Recette) error {
	database.Database.Db.Find(&recette, "id = ?", id)
	if recette.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func GetRecette(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var recette models.Recette
	if err != nil {
		return c.Status(400).JSON("Please ensure that !id is an interger")
	}
	if err := findRecette(id, &recette); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseRecette := CreateResponseRecette(recette)
	return c.Status(200).JSON(responseRecette)
}

func UpdateRecette(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var recette models.Recette
	if err != nil {
		return c.Status(400).JSON("Please ensure that !id is an interger")
	}
	if err := findRecette(id, &recette); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateRecette struct {
		Name         string `json:"name"`
		Descriptions string `json:"descriptions"`
		Ingredients  string `json:"ingredients"`
		Photos       string `json:"photos"`
		Directions   string `json:"directions"`
		Page         string `json:"line"`
		SerialNumber string `json:"serial_number"`
	}
	var updateData UpdateRecette
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	recette.Name = updateData.Name
	recette.Descriptions = updateData.Descriptions
	recette.Ingredients = updateData.Ingredients
	recette.Photos = updateData.Photos
	recette.Directions = updateData.Directions
	recette.Page = updateData.Page
	recette.SerialNumber = updateData.SerialNumber

	database.Database.Db.Save(&recette)

	responseRecette := CreateResponseRecette(recette)
	return c.Status(200).JSON(responseRecette)
}

func DeleteRecette(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var recette models.Recette
	if err != nil {
		return c.Status(400).JSON("Please ensure that !id is an interger")
	}
	if err := findRecette(id, &recette); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := database.Database.Db.Delete(&recette).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).SendString("Successfully Deleted User")
}
