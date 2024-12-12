package handlers

import (
	"github.com/SamstyleGhost/uniblox-assignment/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// When the user first gets on the website, they will have a new id assigned to them which will be stored on the localstorage in frontedn
// This id will be passed here to add the user to our users database
func AddUser(c *fiber.Ctx) error {

	payload := struct {
		ID uuid.UUID `json:"id"`
	}{}

	// Would return error if the request body is not set correctly
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	if err := helpers.AddUserToUsers(payload.ID); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "User added succesfully",
	})

}

// Request : User ID as the url param
// Response : Return the user's cart
func GetUserCart(c *fiber.Ctx) error {
	userParam := c.Params("id")

	userID, err := uuid.Parse(userParam)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
			"message": "ID format invalid",
		})
	}

	user, err := helpers.GetUserCart(userID)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
			"message": "ID format invalid",
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"error":   user,
	})
}

// Request: Item to be added & the user id
func AddItemToCart(c *fiber.Ctx) error {

	payload := struct {
		UserID   uuid.UUID `json:"user_id"`
		ItemID   int       `json:"item_id"`
		Quantity int       `json:"quantity"`
	}{}

	// Would return error if the request body is not set correctly
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
			"message": "Payload structure incorrect",
		})
	}

	cart, err := helpers.AddItemToCart(payload.UserID, payload.ItemID, payload.Quantity)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(201).JSON(&fiber.Map{
		"success": true,
		"cart":    cart,
		"message": "Cart updated successfully",
	})

}
func RemoveItemFromCart(c *fiber.Ctx) {}
func Checkout(c *fiber.Ctx) error {

	payload := struct {
		ID uuid.UUID `json:"user_id"`
	}{}

	// Would return error if the request body is not set correctly
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	user, err := helpers.GetUserCart(payload.ID)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	// Check for discount
	// if discount is applied then return discounted price
	discount, totalPrice := helpers.DiscountCart(user.CartValue)

	// Add to orders.json
	if err := helpers.AddOrder(payload.ID, user.Cart, totalPrice, discount); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	// Empty cart & revise cart_value to 0
	if err := helpers.EmptyCart(payload.ID); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success":        true,
		"cart":           user.Cart,
		"original_price": user.CartValue,
		"discount":       discount,
		"total_price":    totalPrice,
	})
}
