package tests

import (
	"fmt"
	"testing"

	"github.com/SamstyleGhost/uniblox-assignment/helpers"
	"github.com/google/uuid"
)

func TestTraverseItems(t *testing.T) {
	items, err := helpers.TraverseItems()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(items, "", "  ")
	fmt.Print(string(b))
}

func TestFindSelectedItem(t *testing.T) {

	type itemTest struct {
		itemID int
		name   string
	}

	testItems := []itemTest{
		{1, "Office Chair"},
		{2, "Sofa"},
		{3, "Clock"},
		{4, "Round Table"},
		{5, "Knights"},
	}

	for _, testItem := range testItems {
		item, err := helpers.FindSelectedItem(testItem.itemID)
		if err != nil {
			t.Error(err)
		}

		if item.ItemID != testItem.itemID && item.Name != testItem.name {
			t.Error(err)
		}
	}
}

func TestAddUserToUsers(t *testing.T) {

	// The malformed url check is actually done within the handler itself and the helper function is only called if the uuid is valid
	// So, there is no need for the malformed uuid tests
	inputUsers := []string{
		"345e0628-f030-490e-adc1-cf12bdbc1d09",
		// "345e0628-f030-490e-adc1-cf12bdbc1d09sadf",
		// "",
	}

	for _, inputUser := range inputUsers {
		user, _ := uuid.Parse(inputUser)
		if err := helpers.AddUserToUsers(user); err != nil {
			t.Error(err)
		}
	}
}

func TestTraverseUsers(t *testing.T) {
	users, err := helpers.TraverseUsers()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(users, "", "  ")
	fmt.Print(string(b))
}
