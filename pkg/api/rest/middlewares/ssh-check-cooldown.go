package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cloud-barista/cm-beetle/pkg/api/rest/model"
	"github.com/cloud-barista/cm-beetle/pkg/ratelimit"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

const (
	// Mirrors CheckSSHReady's own 30s timeout so only one check per infrastructure runs at a
	// time. Each allowed check opens real SSH connections to every node. It never blocks a
	// client that awaits its response: a not-ready check already runs the full 30s.
	sshCheckMinInterval     = 30 * time.Second
	sshCheckEntryMaxAge     = 10 * time.Minute // Evict an infrastructure's entry after this idle time
	sshCheckCleanupInterval = 5 * time.Minute
)

// SSHCheckCooldown rejects SSH readiness checks that arrive too soon after the previous one for
// the same infrastructure. Counted per (nsId, infraId), because the protected resource is the
// nodes' SSH servers. Attach it to the ssh-ready route only.
func SSHCheckCooldown() echo.MiddlewareFunc {
	cooldown := ratelimit.NewCooldown(sshCheckMinInterval, sshCheckEntryMaxAge, sshCheckCleanupInterval)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			infraRef := c.Param("nsId") + ":" + c.Param("infraId")

			allowed, retryAfter := cooldown.Allow(infraRef)
			if !allowed {
				log.Debug().Msgf("[Middleware] SSH readiness check on cooldown (%s, retry after %v)",
					infraRef, retryAfter.Round(time.Second))

				// An error, not a "not ready" result, so clients back off instead of polling harder.
				c.Response().Header().Set("Retry-After", fmt.Sprintf("%d", ratelimit.RetryAfterSeconds(retryAfter)))
				return c.JSON(http.StatusTooManyRequests, model.SimpleErrorResponse(fmt.Sprintf(
					"SSH readiness check allowed once every %v per infrastructure; retry in %v",
					sshCheckMinInterval, retryAfter.Round(time.Second))))
			}

			return next(c)
		}
	}
}
