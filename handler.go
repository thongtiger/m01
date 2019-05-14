package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// Handlers
func GetProfileHandler(c echo.Context) error {
	var mid = 10 // fixed
	result := mssql.GetProfile(mid)
	return c.JSON(http.StatusOK, result)
}
