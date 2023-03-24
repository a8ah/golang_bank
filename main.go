package main

import (
	"api/database"
	"api/routes"
	"api/utils"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	err := database.ConnectDb()
	if err != nil {
		os.Exit(2)
	}

	app := fiber.New()

	routes.SetUpRouters(app)

	app.Listen(fmt.Sprintf(":%v", utils.Port()))
}
