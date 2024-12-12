package helpers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/SamstyleGhost/uniblox-assignment/models"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
)

func AddUserToUsers(id uuid.UUID) error {

	pathToUsersFile, _ := filepath.Abs("data/users.json") // Doesnt really need the error field here as if there are any errors, it will be handled on the next check
	usersFile, err := os.OpenFile(pathToUsersFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	defer usersFile.Close()

	byteValue, err := io.ReadAll(usersFile)
	if err != nil {
		return err
	}

	var users []models.User
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		return err
	}

	newUser := new(models.User)
	newUser.UserID = id
	newUser.Cart = []models.CartItem{}

	users = append(users, *newUser)

	updatedUsers, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	if err = os.Truncate(pathToUsersFile, 0); err != nil {
		return err
	}

	_, err = usersFile.WriteAt(updatedUsers, 0)
	if err != nil {
		return err
	}

	return nil
}

func TraverseUsers(id uuid.UUID) ([]models.User, error) {

	pathToUsersFile, _ := filepath.Abs("data/users.json") // Doesnt really need the error field here as if there are any errors, it will be handled on the next check
	usersFile, err := os.Open(pathToUsersFile)
	if err != nil {
		return nil, err
	}

	defer usersFile.Close()

	byteValue, err := io.ReadAll(usersFile)
	if err != nil {
		return nil, err
	}

	var users []models.User

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func AddItemToCart(userID uuid.UUID, itemID int, quantity int) ([]models.CartItem, error) {
	pathToUsersFile, _ := filepath.Abs("data/users.json") // Doesnt really need the error field here as if there are any errors, it will be handled on the next check
	usersFile, err := os.OpenFile(pathToUsersFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}

	defer usersFile.Close()

	byteValue, err := io.ReadAll(usersFile)
	if err != nil {
		return nil, err
	}

	var users []models.User

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		return nil, err
	}

	for i, user := range users {
		if user.UserID == userID {

			requestedItem, err := FindSelectedItem(itemID)
			itemToAdd := new(models.CartItem)
			itemToAdd.ItemID = itemID
			itemToAdd.Quantity = quantity

			if err != nil {
				return nil, err
			}

			users[i].Cart = append(users[i].Cart, *itemToAdd)
			users[i].CartValue += requestedItem.Price * float32(quantity)
			updatedUsers, err := json.MarshalIndent(users, "", "  ")
			if err != nil {
				return nil, err
			}

			usersFile.Close()
			usersFile, err = os.OpenFile(pathToUsersFile, os.O_WRONLY|os.O_TRUNC, 0755)
			if err != nil {
				return nil, err
			}
			defer usersFile.Close()

			_, err = usersFile.Write(updatedUsers)
			if err != nil {
				return nil, err
			}
			return users[i].Cart, nil
		}
	}

	return nil, fmt.Errorf("user with ID %s not found", userID)
}

func GetUserCart(userID uuid.UUID) (models.User, error) {

	users, err := TraverseUsers(userID)
	if err != nil {
		return models.User{}, err
	}

	for _, user := range users {
		if user.UserID == userID {
			return user, nil
		}
	}

	// Would return error if user is not found
	return models.User{}, err
}
