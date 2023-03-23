package handlers

import (
	"api/database"
	"api/dto"
	"api/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllClients(c *fiber.Ctx) error {
	clients := []models.Client{}
	db := database.DB.Db

	db.Find(&clients)

	return c.Status(200).JSON(clients)
}

func CreateClients(c *fiber.Ctx) error {
	client := new(models.Client)
	db := database.DB.Db
	if err := c.BodyParser(client); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	db.Create(&client)

	return c.Status(200).JSON(client)
}

func GetClient(c *fiber.Ctx) error {

	id := c.Params("id")
	db := database.DB.Db
	var client models.Client

	db.First(&client, "uuid = ?", id)

	return c.Status(200).JSON(client)
}

func ModifyClients(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")

	// Validating input
	clientInput := new(dto.ClientDTO)
	if err := c.BodyParser(clientInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	var client models.Client
	result := db.First(&client, "uuid = ?", id)

	// validar q la consulta retorne datos, de lo contrario retornar un error
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(
				fiber.Map{
					"message": "No client with that Id Found",
				})
		}
		return c.Status(fiber.StatusBadGateway).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	updates := make(map[string]interface{})
	if clientInput.Name != "" {
		updates["name"] = clientInput.Name
	}
	if clientInput.Surname != "" {
		updates["surname"] = clientInput.Surname
	}
	if clientInput.Dni != "" {
		updates["dni"] = clientInput.Dni
	}

	updates["updated_at"] = time.Now()

	db.Model(&client).Updates(updates)

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   fiber.Map{"client": client},
	})
}
