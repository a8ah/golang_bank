package handlers

import (
	"api/database"
	"api/dto"
	"api/models"

	// "time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetAllTransfer(c *fiber.Ctx) error {
	transactions := []models.Transaction{}
	db := database.DB.Db

	db.Find(&transactions)

	return c.Status(200).JSON(transactions)
}

func GetTransferByAccountId(c *fiber.Ctx) error {

	id := c.Params("id")
	db := database.DB.Db

	transactions := []models.Transaction{}

	db.Where("origin_account = ? OR destination_account >= ?", id, id).Find(&transactions)
	return c.Status(200).JSON(transactions)

}

func GetTransfer(c *fiber.Ctx) error {

	id := c.Params("id")
	db := database.DB.Db
	var transaction models.Transaction

	db.First(&transaction, "uuid = ?", id)

	return c.Status(200).JSON(transaction)
}

func TransferFounds(c *fiber.Ctx) error {

	db := database.DB.Db
	transactionInput := new(dto.TransactionDTO)

	if err := c.BodyParser(transactionInput); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	currencyUUID, err := uuid.Parse(transactionInput.CurrencyUUID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "currency format not alowed. Please  check",
			})
	}

	originAccountUUID, err := uuid.Parse(transactionInput.Origin)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				// "message": err.Error(),
				"message": "origin_accounts format not alowed. Please  check",
			})
	}

	destinationAccountUUID, err := uuid.Parse(transactionInput.Destination)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"message": "destianation_accounts format not alowed. Please  check",
			})
	}

	// originAccount validation
	var originAccount models.Account

	resultOriginAccount := db.First(&originAccount, "uuid = ?", originAccountUUID)
	if err := resultOriginAccount.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(
				fiber.Map{
					"message": "No origin account with that Id Found",
				})
		}
		return c.Status(fiber.StatusBadGateway).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	// destinationAccount validation
	var destinationAccount models.Account

	resultDestinationAccount := db.First(&destinationAccount, "uuid = ?", destinationAccountUUID)
	if err := resultDestinationAccount.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(
				fiber.Map{
					"message": "No destination account with that Id Found",
				})
		}
		return c.Status(fiber.StatusBadGateway).JSON(
			fiber.Map{
				"message": err.Error(),
			})
	}

	// Currency validation
	var currency models.Currency
	resultCurrency := db.First(&currency, "uuid = ?", currencyUUID)

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

	// Validating account has enough founds
	if originAccount.Balance < transactionInput.Amount {
		return c.Status(fiber.StatusNotFound).JSON(
			fiber.Map{
				"message": "Insuficent founds",
			})
	}

	// Validating currancy (same currancy)
	if originAccount.CurrencyUUID != currencyUUID || destinationAccount.CurrencyUUID != currencyUUID {
		return c.Status(fiber.StatusNotFound).JSON(
			fiber.Map{
				"message": "Diferent currancies. Please check",
			})
	}

	transaction := new(models.Transaction)
	transaction.New_origin_account(originAccountUUID)
	transaction.New_destination_account(destinationAccountUUID)
	transaction.New_amount(transactionInput.Amount)
	transaction.New_currency(currencyUUID)

	// transaction.OriginAccount = originAccountUUID
	// transaction.DestinationAccount = destinationAccountUUID
	// transaction.Amount = transactionInput.Amount
	// transaction.CurrencyUUID = currencyUUID

	db.Create(&transaction)

	originUpdateBalance := make(map[string]interface{})
	originUpdateBalance["balance"] = UpdateBalance(originAccount.Balance, transaction.Amount, false)
	db.Model(&originAccount).Updates(originUpdateBalance)

	destinationUpdateBalance := make(map[string]interface{})
	destinationUpdateBalance["balance"] = UpdateBalance(destinationAccount.Balance, transaction.Amount, true)
	db.Model(&destinationAccount).Updates(destinationUpdateBalance)

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"uuid":     transaction.BaseModel.UUID,
			"amount":   transaction.Amount,
			"currency": transaction.CurrencyUUID,
		},
	})

}

func UpdateBalance(initialValue float32, amount float32, credit bool) float32 {
	if credit {
		return initialValue + amount
		// return 0
	} else {
		return initialValue - amount
		// return 1
	}
}
