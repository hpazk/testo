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
		meta := M{"message": "success"}
		return c.JSON(http.StatusOK, meta)
	})

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
