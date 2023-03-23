package handlers

import (
	"api/database"
	"api/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllCurrencies(c *fiber.Ctx) error {
	currencies := []models.Currency{}
	db := database.DB.Db

	db.Find(&currencies)

	return c.Status(200).JSON(currencies)
}

func CreateCurrency(c *fiber.Ctx) error {
	currency := new(models.Currency)
	db := database.DB.Db
	if err := c.BodyParser(currency); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	db.Create(&currency)

	return c.Status(200).JSON(currency)
}

func GetCurrency(c *fiber.Ctx) error {

	id := c.Params("id")
	db := database.DB.Db
	var currency models.Currency

	db.First(&currency, "uuid = ?", id)

	return c.Status(200).JSON(currency)
}

func ModifyCurrency(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")

	// Validating input
	currencyInput := new(models.Currency)
	if err := c.BodyParser(currencyInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	var currency models.Currency
	result := db.First(&currency, "uuid = ?", id)

	// validar q la consulta retorne datos, de lo contrario retornar un error
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(
				fiber.Map{
					"message": "No currency with that Id Found",
				})
		}
		return c.Status(fiber.StatusBadGateway).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	updates := make(map[string]interface{})
	if currencyInput.Name != "" {
		updates["name"] = currencyInput.Name
	}
	if currencyInput.Acronym != "" {
		updates["acronym"] = currencyInput.Acronym
	}

	updates["updated_at"] = time.Now()

	db.Model(&currency).Updates(updates)

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   fiber.Map{"currency": currency},
	})
}
