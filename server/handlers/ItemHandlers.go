package handlers

import (
	"github.com/SamstyleGhost/uniblox-assignment/helpers"
	"github.com/gofiber/fiber/v2"
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
