package main

import (
	"rapid/shoppingcart/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {

	// load template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// static
	app.Static("/public", "./public")

	// controllers
	authController := controllers.InitAuthController()

	app.Get("/login", authController.Login)
	app.Get("/register", authController.Register)
	app.Post("/register", authController.AddRegisteredUser)

	app.Listen(":3000")
}
