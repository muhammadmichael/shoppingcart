package controllers

import (
	"rapid/shoppingcart/database"
	"rapid/shoppingcart/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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

// GET /login
func (controller *AuthController) Login(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}

// GET /register
func (controller *AuthController) Register(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{
		"Title": "Register",
	})
}

// POST /register
func (controller *AuthController) AddRegisteredUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.SendStatus(400) // Bad Request, RegisterForm is not complete
	}

	// Hash password
	bytes, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	sHash := string(bytes)

	// Simpan hashing, bukan plain passwordnya
	user.Password = sHash

	// save product
	err := models.CreateUser(controller.Db, &user)
	if err != nil {
		return c.SendStatus(500) // Server error, gagal menyimpan user
	}
	// if succeed
	return c.Redirect("/login")
}
