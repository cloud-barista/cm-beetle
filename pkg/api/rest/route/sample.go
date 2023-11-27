package route

import (
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/controller"
	"github.com/labstack/echo/v4"
)

// /beetle/sample/*
func RegisterSampleRoutes(g *echo.Group) {
	g.GET("/users", controller.GetUsers)
	g.GET("/users/:id", controller.GetUser)
	g.POST("/users", controller.CreateUser)
	g.PUT("/users/:id", controller.UpdateUser)
	g.DELETE("/users/:id", controller.DeleteUser)
}
