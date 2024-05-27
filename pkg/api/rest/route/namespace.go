package route

import (
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/common"
	"github.com/labstack/echo/v4"
)

// /beetle/ns/*
func RegisterNamespaceRoutes(g *echo.Group) {
	g.POST("", common.RestPostNs)
	g.GET("", common.RestGetAllNs)
	g.GET("/:nsId", common.RestGetNs)
	g.DELETE("/:nsId", common.RestDeleteNs)
}
