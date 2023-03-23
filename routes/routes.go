package routes

import (
	"api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetUpRouters(app *fiber.App) {

	app.Get("/", handlers.Home)

	// Currency endpoints
	app.Get("/currency", handlers.GetAllCurrencies)
	app.Get("/currency/:id", handlers.GetCurrency)
	app.Post("/currency", handlers.CreateCurrency)
	app.Put("/currency/:id", handlers.ModifyCurrency)

	// Client endpoints
	app.Get("/client", handlers.GetAllClients)
	app.Get("/client/:id", handlers.GetClient)
	app.Post("/client", handlers.CreateClients)
	app.Put("/client/:id", handlers.ModifyClients)

	// Account endpoints
	app.Get("/account", handlers.GetAllAccounts)
	app.Get("/account/:id", handlers.GetAccount)
	app.Post("/account", handlers.CreateAccount)
	app.Put("/account/:id/limits", handlers.ModifyLimitsAccount)
	app.Put("/account/:id/secnumb", handlers.ModifySecNumberAccount)
	app.Post("/account/:id/deposit", handlers.DepositAccount)
	app.Post("/account/:id/withdrawals", handlers.WithdrawalsAccount)

	// Transacction
	app.Post("/transfer", handlers.TransferFounds)
	app.Get("/transfer", handlers.GetAllTransfer)
	app.Get("/transfer/:id", handlers.GetTransfer)
	app.Get("/transfer/account/:id", handlers.GetTransferByAccountId)

}
