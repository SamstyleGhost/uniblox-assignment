package main

import (
	"github.com/SamstyleGhost/uniblox-assignment/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.SetRoutes(app)
	app.Listen(":5000")
}
