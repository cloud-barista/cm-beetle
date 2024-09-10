package middlewares

import (
	"context"
	"fmt"
	"time"

	"github.com/cloud-barista/cm-beetle/pkg/logger"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	// "go.opentelemetry.io/otel"
)

// Define Tracing middleware
func TracingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// TODO: Be able to use this code when OpenTelemetry is integrated
		// Start a new tracing context
		// tracer := otel.Tracer("example-tracer")
		// ctx, span := tracer.Start(c.Request().Context(), c.Path())
		// defer span.End()

		// Store trace and span IDs in the context
		// traceId := span.SpanContext().TraceID().String()
		// spanId := span.SpanContext().SpanID().String()

		// Get the context
		ctx := c.Request().Context()

		// [NOTE]
		// For now, we will use the request ID as the trace ID
		// and generate a new span ID for each request
		traceId := c.Response().Header().Get(echo.HeaderXRequestID)
		spanId := fmt.Sprintf("%d", time.Now().UnixNano())

		ctx = context.WithValue(ctx, logger.TraceIdKey, traceId)
		ctx = context.WithValue(ctx, logger.SpanIdKey, spanId)

		// Create a logger with trace_id and span_id and store it in the context
		logger := log.With().Str(string(logger.TraceIdKey), traceId).Str(string(logger.SpanIdKey), spanId).Logger()
		ctx = logger.WithContext(ctx)

		// Set the context in the request
		c.SetRequest(c.Request().WithContext(ctx))

		// [Tracing log] when the request is received
		log.Ctx(ctx).Info().Msg("[tracing] receive request")

		// [Tracing log] before the response is sent
		// Hooks: Before Response
		c.Response().Before(func() {
			// Log the request details
			log.Ctx(ctx).Info().Msg("[tracing] send response")
		})

		// Call the next handler
		return next(c)
	}
}
