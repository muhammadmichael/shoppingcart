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
	prodController := controllers.InitProductController()
	authController := controllers.InitAuthController()

	prod := app.Group("/products")
	prod.Get("/", prodController.GetAllProduct)
	prod.Get("/create", prodController.AddProduct)
	prod.Post("/create", prodController.AddPostedProduct)
	prod.Get("/detail/:id", prodController.DetailProduct)
	prod.Get("/ubah/:id", prodController.UpdateProduct)
	prod.Post("/ubah/:id", prodController.AddUpdatedProduct)
	prod.Get("/hapus/:id", prodController.DeleteProduct)

	app.Get("/login", authController.Login)
	app.Get("/register", authController.Register)
	app.Post("/register", authController.AddRegisteredUser)

	app.Listen(":3000")
}
