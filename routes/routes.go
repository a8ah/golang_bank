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

}
