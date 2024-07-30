package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Pseudo Tumblebug API
	e.GET("/tumblebug/ns", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "GET /tumblebug/ns"})
	})

	e.POST("/tumblebug/ns", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "POST /tumblebug/ns"})
	})

	e.DELETE("/tumblebug/ns/:nsId", func(c echo.Context) error {
		nsId := c.Param("nsId")
		return c.JSON(http.StatusOK, map[string]string{"message": "DELETE /tumblebug/ns/" + nsId})
	})

	e.GET("/tumblebug/ns/:nsId", func(c echo.Context) error {
		nsId := c.Param("nsId")
		return c.JSON(http.StatusOK, map[string]string{"message": "GET /tumblebug/ns/" + nsId})
	})

	e.PUT("/tumblebug/ns/:nsId", func(c echo.Context) error {
		nsId := c.Param("nsId")
		return c.JSON(http.StatusOK, map[string]string{"message": "PUT /tumblebug/ns/" + nsId})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
