package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id       int     `form:"id" json: "id" validate:"required"`
	Name     string  `form:"name" json: "name" validate:"required"`
	Image    string  `form:"image" json: "image" validate:"required"`
	Quantity int     `form:"quantity" json: "quantity" validate:"required"`
	Price    float32 `form:"price" json: "price" validate:"required"`
}

func ReadProducts(db *gorm.DB, products *[]Product) (err error) {
	err = db.Find(products).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadProductById(db *gorm.DB, product *Product, id int) (err error) {
	err = db.Where("id=?", id).First(product).Error
	if err != nil {
		return err
	}
	return nil
}
