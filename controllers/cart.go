package controllers

import (
	"fmt"
	"rapid/shoppingcart/database"
	"rapid/shoppingcart/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CartController struct {
	// Declare variables
	Db *gorm.DB
}

func InitCartController() *CartController {
	db := database.InitDb()
	// gorm sync
	db.AutoMigrate(&models.Cart{})

	return &CartController{Db: db}
}

// GET /addtocart/:cartid/products/:productid
func (controller *CartController) InsertToCart(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	intCartId, _ := strconv.Atoi(params["cartid"])
	intProductId, _ := strconv.Atoi(params["productid"])

	var cart models.Cart
	var product models.Product

	// Find the product first,
	err := models.ReadProductById(controller.Db, &product, intProductId)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	fmt.Println(intCartId)
	// Then find the cart
	errs := models.ReadCartById(controller.Db, &cart, intCartId)
	if errs != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	// Finally, insert the product to cart
	errss := models.InsertProductToCart(controller.Db, &cart, &product)
	if errss != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	return c.Redirect("/products")
}
