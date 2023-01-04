package models

import (
	"goTasks/db"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// GetDB obtiene la conexión a la base de datos
func GetDB() *gorm.DB {
	return db.GetDB()
}

// GetAllProducts obtiene todos los productos de la base de datos
func GetAllProducts() ([]Product, error) {
	var products []Product
	db := GetDB()
	err := db.Find(&products).Error
	return products, err
}

// CreateProduct crea un nuevo producto en la base de datos
func CreateProduct(product *Product) error {
	db := GetDB()
	return db.Create(product).Error
}

// GetProduct obtiene un producto de la base de datos por su ID
func GetProduct(id uint) (*Product, error) {
	var product Product
	db := GetDB()
	err := db.First(&product, id).Error
	return &product, err
}

// UpdateProduct actualiza un producto en la base de datos
func UpdateProduct(product *Product) error {
	db := GetDB()
	return db.Save(product).Error
}

// DeleteProduct elimina un producto de la base de datos
func DeleteProduct(product *Product) error {
	db := GetDB()
	return db.Delete(product).Error
}
