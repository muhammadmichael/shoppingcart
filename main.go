package main

import (
	"rapid/shoppingcart/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html"
)

func main() {
	// session
	store := session.New()

	// load template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// static
	app.Static("/public", "./public")

	// Middleware to check login
	CheckLogin := func(c *fiber.Ctx) error {
		sess, _ := store.Get(c)
		val := sess.Get("username")
		if val != nil {
			return c.Next()
		}

		return c.Redirect("/login")
	}

	// controllers
	prodController := controllers.InitProductController()
	authController := controllers.InitAuthController(store)

	prod := app.Group("/products")
	prod.Get("/", prodController.GetAllProduct)
	prod.Get("/create", CheckLogin, prodController.AddProduct)
	prod.Post("/create", CheckLogin, prodController.AddPostedProduct)
	prod.Get("/detail/:id", prodController.DetailProduct)
	prod.Get("/ubah/:id", CheckLogin, prodController.UpdateProduct)
	prod.Post("/ubah/:id", CheckLogin, prodController.AddUpdatedProduct)
	prod.Get("/hapus/:id", CheckLogin, prodController.DeleteProduct)

	app.Get("/login", authController.Login)
	app.Post("/login", authController.LoginPosted)
	app.Get("/logout", authController.Logout)
	app.Get("/register", authController.Register)
	app.Post("/register", authController.AddRegisteredUser)

	app.Listen(":3000")
}
