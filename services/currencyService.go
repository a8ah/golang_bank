package services

import (
	"api/database"
	"api/models"
	"errors"

	"gorm.io/gorm"
)

func GetAllCurrencies() []models.Currency {
	currencies := []models.Currency{}

	database.DDBB.Find(&currencies)

	return currencies
}

func CreateCurrency(currency *models.Currency) (*models.Currency, error) {

	err := database.DDBB.Create(&currency).Error

	return currency, err
}

func GetCurrency(id string) (models.Currency, error) {

	var currency models.Currency

	err := database.DDBB.First(&currency, "uuid = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return currency, errors.New("no currency found for given id")
	}

	return currency, err
}

func ModifyCurrency(currency models.Currency, updates map[string]interface{}) (models.Currency, error) {

	err := database.DDBB.Model(&currency).Updates(updates).Error

	return currency, err
}
