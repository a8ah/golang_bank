package main

import (
	"api/database"
	"api/routes"
	"api/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnectDb()

	// utils.GetConfig never fail?
	configuration := utils.GetConfig()
	port := configuration.PORT

	if port == "" {
		port = "3000"
	}

	app := fiber.New()

	routes.SetUpRouters(app)

	app.Listen(fmt.Sprintf(":%v", port))
}
