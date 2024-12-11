package handlers

import (
	"fmt"
	"io"
	"os"

	"github.com/SamstyleGhost/uniblox-assignment/models"
	"github.com/gofiber/fiber/v2"

	jsoniter "github.com/json-iterator/go"
)

// Response: List of all items present in the database.
// Havent added pagination
func GetAllItems(c *fiber.Ctx) error {

	items, err := traverseItems()
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

	items, err := traverseItems()
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

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	fmt.Println(payload.ID)

	for _, item := range items {
		if item.ItemID == payload.ID {
			return c.Status(200).JSON(&fiber.Map{
				"success": true,
				"item":    item,
			})
		}
	}

	return c.Status(500).JSON(&fiber.Map{
		"success": false,
		"error":   "Item not found",
	})
}

func AddItemToCart(c *fiber.Ctx) {}
func Checkout(c *fiber.Ctx)      {}

// This is a helper function used to traverse the items.json file
func traverseItems() ([]models.Item, error) {

	itemsFile, err := os.Open("/home/rohan/Work/uniblox-assignment/server/data/items.json")
	if err != nil {
		return nil, err
	}

	defer itemsFile.Close()

	byteValue, err := io.ReadAll(itemsFile)
	if err != nil {
		return nil, err
	}

	var items []models.Item

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(byteValue, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}
