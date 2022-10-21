package controllers

import (
	"fmt"
	"rapid/shoppingcart/database"
	"rapid/shoppingcart/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type ProductController struct {
	// Declare variables
	Db    *gorm.DB
	store *session.Store
}

func InitProductController(s *session.Store) *ProductController {
	db := database.InitDb()
	// gorm sync
	db.AutoMigrate(&models.Product{})

	return &ProductController{Db: db, store: s}
}

// Routing
// GET /products
func (controller *ProductController) GetAllProduct(c *fiber.Ctx) error {
	// Load all Products
	var products []models.Product
	err := models.ReadProducts(controller.Db, &products)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	return c.Render("home", fiber.Map{
		"Title":    "Shopping Cart",
		"Products": products,
	})
}

// GET /products/create
func (controller *ProductController) AddProduct(c *fiber.Ctx) error {
	return c.Render("addproduct", fiber.Map{
		"Title": "Tambah Produk",
	})
}

// POST /products/create
func (controller *ProductController) AddPostedProduct(c *fiber.Ctx) error {
	//myform := new(models.Product)
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Redirect("/products")
	}

	// Parse the multipart form:
	if form, err := c.MultipartForm(); err == nil {
		// => *multipart.Form

		// Get all files from "image" key:
		files := form.File["image"]
		// => []*multipart.FileHeader

		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			// => "tutorial.pdf" 360641 "application/pdf"

			// Save the files to disk:
			product.Image = fmt.Sprintf("public/upload/%s", file.Filename)
			if err := c.SaveFile(file, fmt.Sprintf("public/upload/%s", file.Filename)); err != nil {
				return err
			}
		}
	}

	// save product
	err := models.CreateProduct(controller.Db, &product)
	if err != nil {
		return c.Redirect("/products")
	}
	// if succeed
	return c.Redirect("/products")
}

// GET /products/detail:id
func (controller *ProductController) DetailProduct(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	intId, errs := strconv.Atoi(params["id"])

	if errs != nil {
		fmt.Println(errs)
	}

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, intId)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	sess, err := controller.store.Get(c)
	if err != nil {
		panic(err)
	}
	val := sess.Get("userId")

	return c.Render("productdetail", fiber.Map{
		"Title":   "Detail Product",
		"Product": product,
		"UserId":  val,
	})
}

// GET /products/ubah/:id
func (controller *ProductController) UpdateProduct(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	intId, _ := strconv.Atoi(params["id"])

	var product models.Product
	err := models.ReadProductById(controller.Db, &product, intId)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	return c.Render("productubah", fiber.Map{
		"Title":   "Ubah Produk",
		"Product": product,
	})
}

// POST /products/ubah/:id
func (controller *ProductController) AddUpdatedProduct(c *fiber.Ctx) error {
	var product models.Product

	params := c.AllParams() // "{"id": "1"}"
	intId, _ := strconv.Atoi(params["id"])
	product.Id = intId

	if err := c.BodyParser(&product); err != nil {
		return c.Redirect("/products")
	}

	// Parse the multipart form:
	if form, err := c.MultipartForm(); err == nil {
		// => *multipart.Form

		// Get all files from "image" key:
		files := form.File["image"]
		// => []*multipart.FileHeader

		// Loop through files:
		for _, file := range files {
			fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
			// => "tutorial.pdf" 360641 "application/pdf"

			// Save the files to disk:
			product.Image = fmt.Sprintf("public/upload/%s", file.Filename)
			if err := c.SaveFile(file, fmt.Sprintf("public/upload/%s", file.Filename)); err != nil {
				return err
			}
		}
	}

	// save product
	err := models.UpdateProduct(controller.Db, &product)
	if err != nil {
		return c.Redirect("/products")
	}
	// if succeed
	return c.Redirect("/products")
}

// GET /products/hapus/:id
func (controller *ProductController) DeleteProduct(c *fiber.Ctx) error {
	params := c.AllParams() // "{"id": "1"}"

	intId, errs := strconv.Atoi(params["id"])

	if errs != nil {
		fmt.Println(errs)
	}

	var product models.Product
	err := models.DeleteProductById(controller.Db, &product, intId)
	if err != nil {
		return c.SendStatus(500) // http 500 internal server error
	}

	return c.Redirect("/products")
}
