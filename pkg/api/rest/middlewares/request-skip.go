package middlewares

import (
	"strings"

	"github.com/labstack/echo/v4"
)

// ShouldSkipRequestTracking returns true for management, docs, health, UI, and proxy paths.
func ShouldSkipRequestTracking(c echo.Context) bool {
	p := c.Request().URL.Path
	if p == "" {
		p = c.Path()
	}
	return strings.HasPrefix(p, "/beetle/api") ||
		strings.HasPrefix(p, "/beetle/docs") ||
		strings.HasPrefix(p, "/beetle/swagger") ||
		strings.HasPrefix(p, "/beetle/ui") ||
		p == "/beetle/request" ||
		strings.HasPrefix(p, "/beetle/request/") ||
		p == "/beetle/requests" ||
		strings.HasPrefix(p, "/beetle/requests/") ||
		p == "/beetle/readyz" ||
		p == "/beetle/httpVersion" ||
		strings.HasPrefix(p, "/tumblebug") ||
		strings.HasPrefix(p, "/.well-known") ||
		strings.Contains(p, "/favicon.ico")
}
