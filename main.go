package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type M map[string]interface{}

func main() {
	fmt.Println("running...")

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/users", func(c echo.Context) error {
		meta := M{
			"message":  "success",
			"resource": "users",
		}
		return c.JSON(http.StatusOK, meta)
	})
	e.GET("/products", func(c echo.Context) error {
		meta := M{
			"message":  "success",
			"resource": "products",
		}
		return c.JSON(http.StatusOK, meta)
	})
	e.GET("/orders", func(c echo.Context) error {
		meta := M{
			"message":  "success",
			"resource": "orders",
		}
		return c.JSON(http.StatusOK, meta)
	})

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
