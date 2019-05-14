package main

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var CF Config
var mssql MssqlDB

func init() {
	// initial Config
	CF.GetByENV()
	mssql = MssqlDB{db: NewMssql()}
}
func main() {
	e := echo.New()
	// middleware
	CustomMiddle(e)

	// routes
	e.GET("/", GetProfileHandler)
	e.GET("/test", func(c echo.Context) error {
		if mssql.ExistUsername("burnhee444") {
			// if mssql.ExistUsername("sjtoday") {
			return c.String(200, "ok")
		}
		return c.String(200, "not found.")
	})

	// server start
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(CF.Port)))
}

func CustomMiddle(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST, echo.DELETE},
	}))
}
