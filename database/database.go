package database

import (
	"log"
	"os"

	"github.com/maxime-louis14/api-golang/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var database DbInstance

func ConnectDb() {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(bdGolang)/api-golang"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err.Error())
		os.Exit(2)
	}
	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	// TODO: Add migrations
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	database = DbInstance{Db: db}
}
