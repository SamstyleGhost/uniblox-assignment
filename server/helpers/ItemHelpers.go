package helpers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/SamstyleGhost/uniblox-assignment/models"
	jsoniter "github.com/json-iterator/go"
)

// This is a helper function used to traverse the items.json file
func TraverseItems() ([]models.Item, error) {

	pathToItemsFile, _ := filepath.Abs("data/items.json") // Doesnt really need the error field here as if there are any errors, it will be handled on the next check
	itemsFile, err := os.Open(pathToItemsFile)
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
