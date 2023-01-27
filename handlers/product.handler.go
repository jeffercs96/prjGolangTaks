package handlers

import (
	"goTasks/controllers"

	"github.com/labstack/echo/v4"
)

func ListProducts(c echo.Context) error {
	controllers.ListProducts(c.Response().Writer, c.Request())
	return nil
}
