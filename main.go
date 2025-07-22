package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("GO EXECUTABLE RAN 🚀")
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "GIN SERVER RUNNING SUCCESSFULLY!")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
