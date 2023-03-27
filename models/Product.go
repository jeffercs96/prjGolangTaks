package models

import (
	"goTasks/db"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Code      string    `json:"code" gorm:"unique"`
	Price     uint      `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
func GetProduct(id uuid.UUID) (*Product, error) {
	var product Product
	db := GetDB()
	err := db.Where("id = ?", id).First(&product).Error
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
