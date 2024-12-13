package handlers

import (
	"github.com/SamstyleGhost/uniblox-assignment/helpers"
	"github.com/gofiber/fiber/v2"
)

/*
GET
Response:

	items: []models.Item - List of all items in the database

* Havent added any pagination or other optimizations
*/
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
		"data": fiber.Map{
			"items": items,
		},
	})
}

/*
POST
Request:

	item_id : int - ID of requested item

Response:

	item : model.Item - Requested item
*/
func GetSelectItem(c *fiber.Ctx) error {

	// Get the 'item_id' field from the body of the request
	payload := struct {
		ID int `json:"item_id"`
	}{}

	// Would return error if the request body is not set correctly
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	item, err := helpers.FindSelectedItem(payload.ID)
	if err != nil {
		// Would return error if item is not found
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"data": fiber.Map{
			"item": item,
		},
	})
}
