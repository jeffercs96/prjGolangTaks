package controllers

import (
	"net/http"

	"goTasks/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// GetAllProducts devuelve todos los productos de la base de datos
func GetAllProducts(c echo.Context) error {
	products, err := models.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error getting products")
	}
	return c.JSON(http.StatusOK, products)
}

// CreateProduct crea un nuevo producto en la base de datos
func CreateProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}
	if err := models.CreateProduct(product); err != nil {
		return c.JSON(http.StatusInternalServerError, "Error creating product")
	}
	return c.JSON(http.StatusCreated, product)
}

// GetProduct obtiene un producto de la base de datos por su ID
func GetProduct(c echo.Context) error {
	// Parsear el id del body
	var body struct {
		ID string `json:"id"`
	}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	id, err := uuid.Parse(body.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid product ID")
	}

	// Obtener el producto desde la base de datos
	product, err := models.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}

	// Retornar el producto como respuesta
	return c.JSON(http.StatusOK, product)
}

// UpdateProduct actualiza un producto en la base de datos
func UpdateProduct(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid product ID")
	}
	product, err := models.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request payload")
	}
	if err := models.UpdateProduct(product); err != nil {
		return c.JSON(http.StatusInternalServerError, "Error updating product")
	}
	return c.JSON(http.StatusOK, product)
}

// DeleteProduct elimina un producto de la base de datos
func DeleteProduct(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid product ID")
	}
	product, err := models.GetProduct(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product not found")
	}
	if err := models.DeleteProduct(product); err != nil {
		return c.JSON(http.StatusInternalServerError, "Error deleting product")
	}
	return c.NoContent(http.StatusNoContent)
}
