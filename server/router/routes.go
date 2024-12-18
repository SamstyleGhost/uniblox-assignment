package router

import (
	"github.com/SamstyleGhost/uniblox-assignment/handlers"
	"github.com/SamstyleGhost/uniblox-assignment/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(app *fiber.App) {
	api := app.Group("/api").Group("/v1")

	// Contains apis related to items retrieval
	itemAPI := api.Group("/items")
	itemAPI.Get("/", handlers.GetAllItems)
	itemAPI.Post("/", handlers.GetSelectItem)

	// Contains apis related to user interactions
	// So the endpoint would become '/user/add-to-cart'
	userAPI := api.Group("/user", middleware.CheckUserValidity)
	userAPI.Post("/", handlers.AddUser)
	userAPI.Post("/get-cart", handlers.GetUserCart)
	userAPI.Post("/add-to-cart", handlers.AddItemToCart)
	userAPI.Post("/checkout", handlers.Checkout)
	userAPI.Post("/coupons", handlers.CheckCoupons)
}
