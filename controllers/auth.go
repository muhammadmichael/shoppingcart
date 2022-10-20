package controllers

import (
	"rapid/shoppingcart/database"
	"rapid/shoppingcart/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthController struct {
	// Declare variables
	Db *gorm.DB
}

func InitAuthController() *AuthController {
	db := database.InitDb()
	// gorm sync
	db.AutoMigrate(&models.User{})

	return &AuthController{Db: db}
}

// get /login
func (controller *AuthController) Login(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}

// get /register
func (controller *AuthController) Register(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{
		"Title": "Register",
	})
}
