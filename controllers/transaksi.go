package controllers

import (
	"rapid/shoppingcart/database"
	"rapid/shoppingcart/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type TransaksiController struct {
	// Declare variables
	Db    *gorm.DB
	store *session.Store
}

func InitTransaksiController(s *session.Store) *TransaksiController {
	db := database.InitDb()
	// gorm sync
	db.AutoMigrate(&models.Transaksi{})

	return &TransaksiController{Db: db, store: s}
}

// GET /checkout/:userid
func (controller *TransaksiController) InsertToTransaksi(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	intUserId, _ := strconv.Atoi(params["userid"])

	var transaksi models.Transaksi
	var cart models.Cart

	// Find the product first,
	err := models.ReadAllProductsInCart(controller.Db, &cart, intUserId)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	errs := models.CreateTransaksi(controller.Db, &transaksi, uint(intUserId), cart.Products)
	if errs != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	return c.Redirect("/products")
}

// GET /shoppingcart/:transaksiid
func (controller *TransaksiController) GetTransaksi(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	intTransaksiId, _ := strconv.Atoi(params["transaksiid"])

	var cart models.Cart
	err := models.ReadAllProductsInCart(controller.Db, &cart, intTransaksiId)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	return c.Render("shoppingcart", fiber.Map{
		"Title":    "Detail Product",
		"Products": cart.Products,
	})
}
