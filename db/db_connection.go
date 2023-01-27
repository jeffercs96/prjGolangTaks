package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DBtype     = "mysql"
	DBUser     = "root"
	DBPassword = "jefferson"
	DBName     = "taskPrj"
)

var db *gorm.DB

// Connect establece una conexión con la base de datos MySQL
func Connect() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUser, DBPassword, DBName)
	// Reemplaza "user" y "password" con tus credenciales de acceso
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	return nil
}

// GetDB devuelve la conexión a la base de datos
func GetDB() *gorm.DB {
	return db
}
