package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"api/database"
	"api/utils"
)

func  main()  {

	database.ConnectDb()

	fmt.Println("Dev Configuration\n")
	configuration := utils.GetConfig()
	port := configuration.PORT
	
	if port == "" {
		port = "3000"
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(fmt.Sprintf(":%v", port))
}