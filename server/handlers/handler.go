package handlers

import (
	"github.com/SamstyleGhost/uniblox-assignment/helpers"
	"github.com/SamstyleGhost/uniblox-assignment/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Response: List of all items present in the database.
// Havent added pagination
// ! Right now, I am sending the entire item data in this endpoint, once done, I'll see if I can reduce the data being sent now or adjust the later endpoint in such a way that it only takes in the additional info
func GetAllItems(c *fiber.Ctx) error {

	items, err := helpers.TraverseItems()
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"items":   items,
	})
}

// Request: id - ID of the item which is to be shown
// Response: JSON of the item that has been requested
func GetSelectItem(c *fiber.Ctx) error {

	items, err := helpers.TraverseItems()
	// Would return any error occured while traversing the items.json file
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	// Get the 'id' field from the body of the request
	payload := struct {
		ID int `json:"id"`
	}{}

	// Would return error if the request body is not set correctly
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
			"message": "Incorrect ID",
		})
	}

	for _, item := range items {
		if item.ItemID == payload.ID {
			return c.Status(200).JSON(&fiber.Map{
				"success": true,
				"item":    item,
			})
		}
	}

	// Would return error if item is not found
	return c.Status(400).JSON(&fiber.Map{
		"success": false,
		"error":   "Item not found",
	})
}

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
			"message": "Incorrect ID",
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

func AddItemToCart(c *fiber.Ctx) error {

	payload := struct {
		UserID uuid.UUID   `json:"user_id"`
		Item   models.Item `json:"item"`
	}{}

	// Would return error if the request body is not set correctly
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
			"message": "Payload structure incorrect",
		})
	}

	cart, err := helpers.AddItemToCart(payload.UserID, payload.Item)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
			"message": "Payload structure incorrect",
		})
	}

	return c.Status(201).JSON(&fiber.Map{
		"success": true,
		"cart":    cart,
		"message": "Cart updated successfully",
	})

}
func RemoveItemFromCart(c *fiber.Ctx) {}
func Checkout(c *fiber.Ctx)           {}
