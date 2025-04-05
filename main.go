package main

import (
	"learngithubactions/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/add-strings", func(c echo.Context) error {
		a := c.QueryParam("a")
		b := c.QueryParam("b")

		r, err := utils.AddStrings(a, b)
		if err != nil {
			return err
		}

		return c.String(http.StatusOK, r)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
