package helpers

import (
	"fmt"

	jsonfileworker "github.com/SamstyleGhost/uniblox-assignment/lib/json-file-worker"
	"github.com/SamstyleGhost/uniblox-assignment/models"
)

var (
	itemsPath = "data/items.json"
)

// This is a helper function used to traverse the items.json file
func TraverseItems() ([]models.Item, error) {

	var items []models.Item
	err := jsonfileworker.GetAllObjects(itemsPath, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func FindSelectedItem(itemID int) (models.Item, error) {

	items, err := TraverseItems()
	if err != nil {
		return models.Item{}, err
	}

	for _, item := range items {
		if item.ItemID == itemID {
			return item, nil
		}
	}

	return models.Item{}, fmt.Errorf("item with itemID %d not found", itemID)
}
