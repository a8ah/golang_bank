package handlers

import (
	"api/database"
	"api/dto"
	"api/models"
	"math/rand"

	// "time"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllAccounts(c *fiber.Ctx) error {
	accounts := []models.Account{}
	db := database.DB.Db

	db.Find(&accounts)

	return c.Status(200).JSON(accounts)
}

func CreateAccount(c *fiber.Ctx) error {

	db := database.DB.Db
	accountInput := new(dto.AccountCreateDTO)

	if err := c.BodyParser(accountInput); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	// client validation
	var client models.Client

	resultClient := db.First(&client, "uuid = ?", accountInput.ClientUUID)
	if err := resultClient.Error; err != nil {
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

	// Currency validation
	var currency models.Currency
	resultCurrency := db.First(&currency, "uuid = ?", accountInput.CurrencyUUID)
	// validar q la consulta retorne datos, de lo contrario retornar un error
	if err := resultCurrency.Error; err != nil {
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

	var accountNumber uint64
	fmt.Sscan(accountInput.Number, &accountNumber)
	var min, max = 1000, 9999
	var secureNumber = rand.Intn(max)
	if secureNumber < min {
		secureNumber += min
	}

	account := new(models.Account)
	account.ClientUUID = client.BaseModel.UUID
	account.CurrencyUUID = currency.BaseModel.UUID
	account.Number = accountNumber
	account.SecureNumber = uint16(secureNumber)

	db.Create(&account)

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"uuid":           account.BaseModel.UUID,
			"account_number": account.Number,
			"currency":       account.CurrencyUUID,
		},
	})
}

func GetAccount(c *fiber.Ctx) error {

	id := c.Params("id")
	db := database.DB.Db
	var account models.Account

	db.First(&account, "uuid = ?", id)

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"uuid":           account.BaseModel.UUID,
			"account_number": account.Number,
			"currency":       account.CurrencyUUID,
			"balance":        account.Balance,
			"limit":          account.Limit,
		},
	})
}

func ModifyLimitsAccount(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	newAccountValues := new(models.Account)
	var account models.Account

	// Validating input

	if err := c.BodyParser(newAccountValues); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	result := db.First(&account, "uuid = ?", id)
	// validar q la consulta retorne datos, de lo contrario retornar un error
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(
				fiber.Map{
					"message": "No account with that Id Found",
				})
		}
		return c.Status(fiber.StatusBadGateway).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	updates := make(map[string]interface{})
	updates["limit"] = newAccountValues.Limit

	db.Model(&account).Updates(updates)

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"uuid":           account.BaseModel.UUID,
			"account_number": account.Number,
			"currency":       account.CurrencyUUID,
			"limit":          account.Limit,
		},
	})
}

func ModifySecNumberAccount(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	newAccountValues := new(models.Account)
	var account models.Account

	// Validating input
	if err := c.BodyParser(newAccountValues); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	result := db.First(&account, "uuid = ?", id)
	// validar q la consulta retorne datos, de lo contrario retornar un error
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(
				fiber.Map{
					"message": "No account with that Id Found",
				})
		}
		return c.Status(fiber.StatusBadGateway).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	updates := make(map[string]interface{})
	updates["secure_number"] = newAccountValues.SecureNumber

	db.Model(&account).Updates(updates)

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"uuid":           account.BaseModel.UUID,
			"account_number": account.Number,
			"currency":       account.CurrencyUUID,
			"limit":          account.Limit,
		},
	})
}

func DepositAccount(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	inputAccount := new(models.Account)
	var account models.Account

	// Validating input
	if err := c.BodyParser(inputAccount); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	result := db.First(&account, "uuid = ?", id)
	// validar q la consulta retorne datos, de lo contrario retornar un error
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(
				fiber.Map{
					"message": "No account with that Id Found",
				})
		}
		return c.Status(fiber.StatusBadGateway).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	updates := make(map[string]interface{})
	updates["balance"] = UpdateBalance(account.Balance, inputAccount.Balance, true)

	db.Model(&account).Updates(updates)

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"uuid":           account.BaseModel.UUID,
			"account_number": account.Number,
			"currency":       account.CurrencyUUID,
			"limit":          account.Limit,
			"balance":        account.Balance,
		},
	})
}

func WithdrawalsAccount(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	inputAccount := new(models.Account)
	var account models.Account

	// Validating input
	if err := c.BodyParser(inputAccount); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	result := db.First(&account, "uuid = ?", id)
	// validar q la consulta retorne datos, de lo contrario retornar un error
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(
				fiber.Map{
					"message": "No account with that Id Found",
				})
		}
		return c.Status(fiber.StatusBadGateway).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	// Validating account has enough founds
	if account.Balance < inputAccount.Balance {
		return c.Status(fiber.StatusNotFound).JSON(
			fiber.Map{
				"message": "Insuficent founds",
			})
	}

	updates := make(map[string]interface{})
	updates["balance"] = UpdateBalance(account.Balance, inputAccount.Balance, false)

	db.Model(&account).Updates(updates)

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"uuid":           account.BaseModel.UUID,
			"account_number": account.Number,
			"currency":       account.CurrencyUUID,
			"limit":          account.Limit,
			"balance":        account.Balance,
		},
	})
}
