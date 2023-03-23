package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"api/database"
	"api/utils"
	"api/routes"
)

func  main()  {

	database.ConnectDb()

	configuration := utils.GetConfig()
	port := configuration.PORT

	if port == "" {
		port = "3000"
	}

	app := fiber.New()

	routes.SetUpRouters(app)

	app.Listen(fmt.Sprintf(":%v", port))
}