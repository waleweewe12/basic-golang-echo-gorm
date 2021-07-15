package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/:name", func(c echo.Context) error {
		resp := fmt.Sprintf("Hello %s", c.Param("name"))
		return c.String(http.StatusOK, resp)
	})
	e.POST("/", func(c echo.Context) error {
		var request struct {
			Name string `json:"name"`
		}
		if err := c.Bind(&request); err != nil {
			return err
		}
		resp := fmt.Sprintf("Hello %s", request.Name)
		return c.String(http.StatusOK, resp)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
