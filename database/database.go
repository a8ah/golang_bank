package database

import (
	"api/utils"
	"log"
	"os"
	"fmt"
	// "api/models"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/logger"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	configuration := utils.GetConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		configuration.DB_HOST,
		configuration.DB_USERNAME,
		configuration.DB_PASSWORD,
		configuration.DB_NAME,
		configuration.DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to daabase. \n", err)
		os.Exit(2)
	}

	log.Print("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	// log.Println("running migrations")
	// db.AutoMigrate(&models.Client{})
	// db.AutoMigrate(&models.Account{})

	DB = Dbinstance{
		Db: db,
	}
}
