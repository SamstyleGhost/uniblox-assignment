package router

import (
	"github.com/SamstyleGhost/uniblox-assignment/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	api := app.Group("/api").Group("v1")

	// Contains apis related to items retrieval
	itemAPI := api.Group("/items")
	itemAPI.Get("/", handlers.GetAllItems)
	itemAPI.Post("/", handlers.GetSelectItem)

	// Contains apis related to user interactions
	userAPI := api.Group("/user")
	userAPI.Post("/", handlers.AddUser)
	userAPI.Post("/add-to-cart", handlers.AddItemToCart)
	userAPI.Get("/:id", handlers.GetUserCart)
	userAPI.Post("/checkout", handlers.Checkout)
}
