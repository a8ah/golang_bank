package routes

import (
	"api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetUpRouters(app *fiber.App) {

	app.Get("/", handlers.Home)

	// Client endpoints
	app.Get("/client", handlers.GetAllClients)
	app.Get("/client/:id", handlers.GetClient)
	app.Post("/client", handlers.CreateClients)
	app.Put("/client/:id", handlers.ModifyClients)
}
