package database

import (
	"api/utils"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Dbinstance struct {
	Db *gorm.DB
}

// Could you propose a better way to share this object?
var DB Dbinstance

// To make an entire app fail from an internal package is not a good practice.
// Can you modify this methode in a way that the code using it is notified of
// the error in case it happens?
func ConnectDb() error {
	configuration, err := utils.GetConfig()
	if err != nil {
		return err
	}

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
		log.Print("Failed to connect to database. \n", err)
		return err
	}

	log.Print("Database Connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	// log.Println("running migrations")
	// db.AutoMigrate(&models.Currency{})
	// db.AutoMigrate(&models.Client{})
	// db.AutoMigrate(&models.Account{})
	// db.AutoMigrate(&models.Transaction{})

	DB = Dbinstance{
		Db: db,
	}

	return nil
}
