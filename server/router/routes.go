package router

import (
	"github.com/SamstyleGhost/uniblox-assignment/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	api := app.Group("/api").Group("v1")

	api.Get("/items", handlers.GetAllItems)
	api.Post("/items", handlers.GetSelectItem)
}
