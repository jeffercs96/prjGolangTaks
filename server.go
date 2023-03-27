package main

import (
	"goTasks/controllers"
	"goTasks/db"
	"goTasks/models"
	"log"
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Inicializar el router

	// Endpoints para los productos

	// Routes
	e.GET("/", hello)
	// e.GET("/employees", controllers.GetEmployees)
	// Start server
	// Crea 5 productos

	// Rutas para GetAllProducts
	e.GET("/products", controllers.GetAllProducts)

	// Rutas para CreateProduct
	e.POST("/createproduct", controllers.CreateProduct)

	// Rutas para GetProduct, UpdateProduct y DeleteProduct
	e.POST("/getproduct", controllers.GetProduct)
	e.PUT("/updateproduct", controllers.UpdateProduct)
	e.DELETE("/deleteproduct", controllers.DeleteProduct)
	/* o tambien e.GET("/products", echo.WrapHandler(http.HandlerFunc(controllers.ListProducts))) */
	// Crea la tabla "products" en la base de datos si no existe
	// models.CreateProduct(&models.Product{Code: "P1", Price: 100})
	// models.CreateProduct(&models.Product{Code: "P2", Price: 200})
	// models.CreateProduct(&models.Product{Code: "P3", Price: 300})
	// models.CreateProduct(&models.Product{Code: "P4", Price: 400})
	// models.CreateProduct(&models.Product{Code: "P5", Price: 500})
	db := db.GetDB()
	db.AutoMigrate(&models.Product{})
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func createInitialData() {
	// Crea algunos productos de ejemplo
	products := []*models.Product{
		{Code: "P1", Price: 100},
		{Code: "P2", Price: 200},
		{Code: "P3", Price: 300},
		{Code: "P4", Price: 400},
		{Code: "P5", Price: 500},
	}
	for _, product := range products {
		if err := models.CreateProduct(product); err != nil {
			log.Fatalf("Error creando producto: %v", err)
		}
	}
}
