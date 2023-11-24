package route

import (
	"github.com/cloud-barista/cm-beetle/pkg/api/rest/controller"
	"github.com/labstack/echo/v4"
)

// /beetle/migration/*
func RegisterMigrationRoutes(g *echo.Group) {
	g.POST("/infra", controller.MigrateInfra)
	g.POST("/infra/network", controller.MigrateInfra)
	g.POST("/infra/storage", controller.MigrateInfra)
	g.POST("/infra/instance", controller.MigrateInfra)
}
