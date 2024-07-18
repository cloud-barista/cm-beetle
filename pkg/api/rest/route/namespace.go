package route

import (
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/controller"
	"github.com/labstack/echo/v4"
)

// /beetle/ns/*
func RegisterNamespaceRoutes(g *echo.Group) {
	g.POST("", controller.RestPostNs)
	g.GET("", controller.RestGetAllNs)
	g.GET("/:nsId", controller.RestGetNs)
	g.DELETE("/:nsId", controller.RestDeleteNs)
}
