package handlers

import (
	"api/models"
	"api/services"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllCurrencies(c *fiber.Ctx) error {
	currencies := services.GetAllCurrencies()

	return c.Status(200).JSON(currencies)
}

func CreateCurrency(c *fiber.Ctx) error {
	currency := new(models.Currency)

	if err := c.BodyParser(currency); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	result, err := services.CreateCurrency(currency)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	return c.Status(200).JSON(result)
}

func GetCurrency(c *fiber.Ctx) error {

	id := c.Params("id")
	currency, err := services.GetCurrency(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}
	return c.Status(200).JSON(currency)
}

func ModifyCurrency(c *fiber.Ctx) error {

	id := c.Params("id")

	// Validating input
	currencyInput := new(models.Currency)
	if err := c.BodyParser(currencyInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	currency, err := services.GetCurrency(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
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

	result, err := services.ModifyCurrency(currency, updates)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   fiber.Map{"currency": result},
	})
}
