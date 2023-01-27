package main

import (
	"goTasks/db"
	"goTasks/handlers"
	"goTasks/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	db.Connect()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Inicializar el router

	// Endpoints para los productos

	// Routes
	e.GET("/", hello)
	// e.GET("/employees", controllers.GetEmployees)
	// Start server
	// Crea 5 productos
	db := db.GetDB()
	db.AutoMigrate(&models.Product{})
	e.GET("/products", handlers.ListProducts)
	/* o tambien e.GET("/products", echo.WrapHandler(http.HandlerFunc(controllers.ListProducts))) */
	// Crea la tabla "products" en la base de datos si no existe
	// models.CreateProduct(&models.Product{Code: "P1", Price: 100})
	// models.CreateProduct(&models.Product{Code: "P2", Price: 200})
	// models.CreateProduct(&models.Product{Code: "P3", Price: 300})
	// models.CreateProduct(&models.Product{Code: "P4", Price: 400})
	// models.CreateProduct(&models.Product{Code: "P5", Price: 500})
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
