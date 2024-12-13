package handlers

import (
	"github.com/SamstyleGhost/uniblox-assignment/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

/*
POST
Request:

	user_id: uuid
*/
func AddUser(c *fiber.Ctx) error {

	payload := struct {
		ID uuid.UUID `json:"user_id"`
	}{}

	// Would return error if the request body is not set correctly
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	if payload.ID == uuid.Nil {
		return c.Status(400).JSON(&fiber.Map{
			"error": "Nil uuid received",
		})
	}

	if err := helpers.AddUserToUsers(payload.ID); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"message": "User added succesfully",
	})
}

/*
GET
Parameters:

	user-id: uuid

Response:

	user: models.User
*/
func GetUserCart(c *fiber.Ctx) error {

	userParam := c.Params("id")

	userID, err := uuid.Parse(userParam)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := helpers.GetUserCart(userID)
	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"user": user,
	})
}

/*
POST
Request:

	user_id: uuid
	item_id: int
	quantity: int

Response:

	cart: []models.CartItem - Updated cart
*/
func AddItemToCart(c *fiber.Ctx) error {

	payload := struct {
		UserID   uuid.UUID `json:"user_id"`
		ItemID   int       `json:"item_id"`
		Quantity int       `json:"quantity"`
	}{}

	// Would return error if the request body is not set correctly
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"error":   err.Error(),
			"message": "Payload structure incorrect",
		})
	}

	// If the quantity is 0, then do not add to cart
	if payload.Quantity == 0 {
		return c.Status(400).JSON(&fiber.Map{
			"error": "No quantity selected",
		})
	}

	cart, err := helpers.AddItemToCart(payload.UserID, payload.ItemID, payload.Quantity)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(&fiber.Map{
		"cart": cart,
	})
}

/*
POST
Request:

	user_id: uuid - ID of the user
	coupon_code: uuid - Coupon code for discount (Optional - defaults to uuid.Nil)

Response:

	cart: []models.CartItem - The items bought
	original_price: float32 - The original price of the cart without any discounts
	discount: float32 - The amount of discount given
	total_price: float32 - Original price - discount
	coupon: uuid - New coupon code if available
*/
func Checkout(c *fiber.Ctx) error {

	payload := struct {
		ID     uuid.UUID `json:"user_id"`
		Coupon uuid.UUID `json:"coupon_code"`
	}{}

	// Would return error if the request body is not set correctly
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := helpers.GetUserCart(payload.ID)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	// Checks the condition when cart is empty. No checkout possible then
	if len(user.Cart) == 0 {
		return c.Status(400).JSON(&fiber.Map{
			"error": "cart empty",
		})
	}

	var discount bool

	// Validate coupon if provided
	if payload.Coupon != uuid.Nil {
		discount, _ = helpers.ValidateCoupon(payload.Coupon, payload.ID)
	}

	// Give a new coupon every 3rd checkout
	newCoupon, _ := helpers.GenerateCoupon(payload.ID)

	var discountPrice float32
	totalPrice := user.CartValue

	if discount {
		discountPrice = user.CartValue * 0.1 // 10% discount if coupon is valid
		totalPrice -= discountPrice
	}

	// Add to orders.json
	if err := helpers.AddOrder(payload.ID, user.Cart, totalPrice, discountPrice); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	// Empty cart & revise cart_value to 0
	if err := helpers.EmptyCart(payload.ID); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	// TODO: Subtract the quantity of items bought from the overall stock

	return c.Status(200).JSON(&fiber.Map{
		"cart":           user.Cart,
		"original_price": user.CartValue,
		"discount":       discountPrice,
		"total_price":    totalPrice,
		"coupon":         newCoupon,
	})
}

func CheckCoupons(c *fiber.Ctx) error {
	payload := struct {
		UserID uuid.UUID `json:"user_id"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	userCoupons, err := helpers.CheckCoupons(payload.UserID)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(userCoupons)
}
