package helpers

import (
	"fmt"

	"github.com/google/uuid"

	jsonfileworker "github.com/SamstyleGhost/uniblox-assignment/lib/json-file-worker"
	"github.com/SamstyleGhost/uniblox-assignment/models"
)

// A counter variable to count the number of orders
// Another way was to count the len of the array in orders.json
var (
	counter         = 0
	usersFilePath   = "data/users.json"
	ordersFilePath  = "data/orders.json"
	couponsFilePath = "data/coupons.json"
)

// The id received to this function is guaranteed to be a valid uuid. The check for validity has been done in the handler itself
func AddUserToUsers(id uuid.UUID) error {

	var users []models.User
	err := jsonfileworker.GetAllObjects(usersFilePath, &users)
	if err != nil {
		return err
	}

	newUser := new(models.User)
	newUser.UserID = id
	newUser.Cart = []models.CartItem{}

	users = append(users, *newUser)

	err = jsonfileworker.SetAllObjects(usersFilePath, users)
	if err != nil {
		return err
	}

	return nil
}

func TraverseUsers() ([]models.User, error) {

	var users []models.User
	err := jsonfileworker.GetAllObjects(usersFilePath, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func AddItemToCart(userID uuid.UUID, itemID int, quantity int) ([]models.CartItem, error) {

	var users []models.User

	err := jsonfileworker.GetAllObjects(usersFilePath, &users)
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

			err = jsonfileworker.SetAllObjects(usersFilePath, users)

			if err != nil {
				return nil, err
			}
			return users[i].Cart, nil
		}
	}

	return nil, fmt.Errorf("user with ID %s not found", userID)
}

func GetUserCart(userID uuid.UUID) (models.User, error) {

	users, err := TraverseUsers()
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

	var orders []models.Order

	err := jsonfileworker.GetAllObjects(ordersFilePath, &orders)
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

	err = jsonfileworker.SetAllObjects(ordersFilePath, orders)
	if err != nil {
		return err
	}

	counter++

	return nil
}

func EmptyCart(userID uuid.UUID) error {

	var users []models.User

	if err := jsonfileworker.GetAllObjects(usersFilePath, &users); err != nil {
		return err
	}

	for i, user := range users {
		if user.UserID == userID {

			users[i].Cart = []models.CartItem{}
			users[i].CartValue = 0

			if err := jsonfileworker.SetAllObjects(usersFilePath, &users); err != nil {
				return err
			}

			return nil
		}
	}

	return fmt.Errorf("user with ID %s not found", userID)
}

func GenerateCoupon(userID uuid.UUID) (uuid.UUID, error) {

	// Coupons will only be generated for every 3rd order
	if counter%3 == 0 {
		coupon := uuid.New()
		var coupons []models.Coupon

		if err := jsonfileworker.GetAllObjects(couponsFilePath, &coupons); err != nil {
			return uuid.Nil, err
		}

		newCoupon := new(models.Coupon)
		newCoupon.UserID = userID
		newCoupon.CouponCode = coupon

		coupons = append(coupons, *newCoupon)

		if err := jsonfileworker.SetAllObjects(couponsFilePath, &coupons); err != nil {
			return uuid.Nil, err
		}

		return coupon, nil
	}

	return uuid.Nil, fmt.Errorf("no coupon for this order")
}

func ValidateCoupon(coupon uuid.UUID, userID uuid.UUID) (bool, error) {

	var coupons []models.Coupon

	if err := jsonfileworker.GetAllObjects(couponsFilePath, &coupons); err != nil {
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

	if err := jsonfileworker.SetAllObjects(couponsFilePath, &coupons); err != nil {
		return false, err
	}

	return true, nil
}

func CheckCoupons(userID uuid.UUID) ([]models.Coupon, error) {

	var coupons []models.Coupon
	if err := jsonfileworker.GetAllObjects(couponsFilePath, &coupons); err != nil {
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
