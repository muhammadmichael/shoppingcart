package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID   uint
	Products []*Product `gorm:"many2many:cart_products;"`
}

// CRUD

func CreateCart(db *gorm.DB, newCart *Cart) (err error) {
	err = db.Create(newCart).Error
	if err != nil {
		return err
	}
	return nil
}

func InsertProductToCart(db *gorm.DB, product *Product) (err error) {
	err = db.Create(product).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadCarts(db *gorm.DB, products *[]Product) (err error) {
	err = db.Find(products).Error
	if err != nil {
		return err
	}
	return nil
}
