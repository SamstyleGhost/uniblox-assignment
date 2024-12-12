package router

import (
	"github.com/SamstyleGhost/uniblox-assignment/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	api := app.Group("/api").Group("v1")

	api.Get("/items", handlers.GetAllItems)
	api.Post("/items", handlers.GetSelectItem)

	api.Post("/user", handlers.AddUser)
	api.Post("/user/add-to-cart", handlers.AddItemToCart)
	api.Get("/user/:id", handlers.GetUserCart)
	api.Post("/user/checkout", handlers.Checkout)
}
