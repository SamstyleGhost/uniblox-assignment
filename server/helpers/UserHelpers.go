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

// A counter variable to count the number of orders
// Another way was to count the len of the array in orders.json
var counter = 0

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

func AddOrder(userID uuid.UUID, cartItems []models.CartItem, orderValue float32, discount float32) error {

	pathToOrdersFile, _ := filepath.Abs("data/orders.json") // Doesnt really need the error field here as if there are any errors, it will be handled on the next check
	ordersFile, err := os.OpenFile(pathToOrdersFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	defer ordersFile.Close()

	byteValue, err := io.ReadAll(ordersFile)
	if err != nil {
		return err
	}

	var orders []models.Order
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err = json.Unmarshal(byteValue, &orders)
	if err != nil {
		return err
	}

	newOrder := new(models.Order)
	newOrder.OrderID = counter
	newOrder.UserID = userID
	newOrder.Items = cartItems
	newOrder.OrderValue = orderValue
	newOrder.Discount = discount

	orders = append(orders, *newOrder)

	updatedOrders, err := json.MarshalIndent(orders, "", "  ")
	if err != nil {
		return err
	}

	if err = os.Truncate(pathToOrdersFile, 0); err != nil {
		return err
	}

	_, err = ordersFile.WriteAt(updatedOrders, 0)
	if err != nil {
		return err
	}

	counter++

	return nil
}

func EmptyCart(userID uuid.UUID) error {

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

	for i, user := range users {
		if user.UserID == userID {

			users[i].Cart = []models.CartItem{}
			users[i].CartValue = 0
			updatedUsers, err := json.MarshalIndent(users, "", "  ")
			if err != nil {
				return err
			}

			usersFile.Close()
			usersFile, err = os.OpenFile(pathToUsersFile, os.O_WRONLY|os.O_TRUNC, 0755)
			if err != nil {
				return err
			}
			defer usersFile.Close()

			_, err = usersFile.Write(updatedUsers)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return fmt.Errorf("user with ID %s not found", userID)
}

func GenerateCoupon(userID uuid.UUID) (uuid.UUID, error) {
	if counter%3 == 0 {
		coupon := uuid.New()

		pathToCouponsFile, _ := filepath.Abs("data/coupons.json") // Doesnt really need the error field here as if there are any errors, it will be handled on the next check
		couponsFile, err := os.OpenFile(pathToCouponsFile, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return uuid.Nil, err
		}

		defer couponsFile.Close()

		byteValue, err := io.ReadAll(couponsFile)
		if err != nil {
			return uuid.Nil, err
		}

		var coupons []models.Coupon
		var json = jsoniter.ConfigCompatibleWithStandardLibrary

		err = json.Unmarshal(byteValue, &coupons)
		if err != nil {
			return uuid.Nil, err
		}

		newCoupon := new(models.Coupon)
		newCoupon.UserID = userID
		newCoupon.CouponCode = coupon

		coupons = append(coupons, *newCoupon)
		updatedCoupons, err := json.MarshalIndent(coupons, "", "  ")
		if err != nil {
			return uuid.Nil, err
		}

		couponsFile.Close()
		couponsFile, err = os.OpenFile(pathToCouponsFile, os.O_WRONLY|os.O_TRUNC, 0755)
		if err != nil {
			return uuid.Nil, err
		}
		defer couponsFile.Close()

		_, err = couponsFile.Write(updatedCoupons)
		if err != nil {
			return uuid.Nil, err
		}

		return coupon, nil
	}

	return uuid.Nil, fmt.Errorf("no coupon for this order")
}

func ValidateCoupon(coupon uuid.UUID, userID uuid.UUID) (bool, error) {

	pathToCouponsFile, _ := filepath.Abs("data/coupons.json") // Doesnt really need the error field here as if there are any errors, it will be handled on the next check
	couponsFile, err := os.OpenFile(pathToCouponsFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return false, err
	}

	defer couponsFile.Close()

	byteValue, err := io.ReadAll(couponsFile)
	if err != nil {
		return false, err
	}

	var coupons []models.Coupon
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	err = json.Unmarshal(byteValue, &coupons)
	if err != nil {
		return false, err
	}

	index := -1
	for i, cpn := range coupons {
		if cpn.CouponCode == coupon && cpn.UserID == userID {
			index = i
			break
		}
	}

	if index == -1 {
		return false, fmt.Errorf("coupon code %s invalid", coupon)
	}

	coupons = append(coupons[:index], coupons[index+1:]...)

	updatedCoupons, err := json.MarshalIndent(coupons, "", "  ")
	if err != nil {
		return false, err
	}

	couponsFile.Close()
	couponsFile, err = os.OpenFile(pathToCouponsFile, os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return false, err
	}
	defer couponsFile.Close()

	_, err = couponsFile.Write(updatedCoupons)
	if err != nil {
		return false, err
	}

	return true, nil
}

func CheckCoupons(userID uuid.UUID) ([]models.Coupon, error) {

	pathToCouponsFile, _ := filepath.Abs("data/coupons.json") // Doesnt really need the error field here as if there are any errors, it will be handled on the next check
	couponsFile, err := os.Open(pathToCouponsFile)
	if err != nil {
		return nil, err
	}

	defer couponsFile.Close()

	byteValue, err := io.ReadAll(couponsFile)
	if err != nil {
		return nil, err
	}

	var coupons []models.Coupon
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	err = json.Unmarshal(byteValue, &coupons)
	if err != nil {
		return nil, err
	}

	userCoupons := []models.Coupon{}

	for _, cpn := range coupons {
		if cpn.UserID == userID {
			userCoupons = append(userCoupons, cpn)
		}
	}

	return userCoupons, nil
}
