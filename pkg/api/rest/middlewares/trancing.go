package middlewares

import (
	"context"
	"fmt"
	"strings"
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

		var traceId, spanId string

		// [OpenTelemetry]
		// Uncomment the following code when OTel is integrated
		/*
			span := trace.SpanFromContext(ctx)
			if span.SpanContext().IsValid() {
				traceId = span.SpanContext().TraceID().String()
				spanId = span.SpanContext().SpanID().String()
			}
		*/

		// [Fallback / Current Implementation]
		// TODO: Remove this fallback logic when OpenTelemetry is fully integrated.
		// If OTel IDs are not available, try to use/generate traceparent manually.
		if traceId == "" {
			// 1. Try to parse traceparent header manually
			tp := c.Request().Header.Get("traceparent")
			if tp == "" {
				// Generate dummy traceparent if missing
				// Format: 00-{traceId(32)}-{spanId(16)}-01
				tid := fmt.Sprintf("%016x%016x", time.Now().UnixNano(), time.Now().UnixNano())
				pid := fmt.Sprintf("%016x", time.Now().UnixNano())
				tp = fmt.Sprintf("00-%s-%s-01", tid, pid)
				c.Request().Header.Set("traceparent", tp)
			}

			// Parse traceId from traceparent
			parts := strings.Split(tp, "-")
			if len(parts) == 4 {
				traceId = parts[1]
				// Generate a new SpanID for the current span (16 hex chars)
				spanId = fmt.Sprintf("%016x", time.Now().UnixNano())
			} else {
				// Fallback if parsing failed (should not happen if we generated it)
				traceId = c.Response().Header().Get(echo.HeaderXRequestID)
				if traceId == "" {
					traceId = fmt.Sprintf("trace-%d", time.Now().UnixNano())
				}
				spanId = fmt.Sprintf("%d", time.Now().UnixNano())
			}
		}

		// Store trace and span IDs in the context
		if traceId != "" {
			ctx = context.WithValue(ctx, logger.TraceIdKey, traceId)
		}
		if spanId != "" {
			ctx = context.WithValue(ctx, logger.SpanIdKey, spanId)
		}

		// [Fallback / Current Implementation]
		// TODO: Remove this block when OpenTelemetry is fully integrated.
		// Currently, we bake the IDs into the logger instance so they appear in all logs derived from this context.
		// When OTel is integrated, the TracingHook will extract IDs from the Context dynamically,
		// so we will not need to bake them into the logger here.
		loggerContext := log.With()
		if traceId != "" {
			loggerContext = loggerContext.Str(string(logger.TraceIdKey), traceId)
		}
		if spanId != "" {
			loggerContext = loggerContext.Str(string(logger.SpanIdKey), spanId)
		}
		childLogger := loggerContext.Logger()
		ctx = childLogger.WithContext(ctx)

		// [OpenTelemetry]
		// Uncomment the following line when OTel is integrated (and remove the Fallback block above).
		// This ensures the global logger (with TracingHook) is associated with the context.
		// ctx = log.Logger.WithContext(ctx)

		// Set the context in the request
		c.SetRequest(c.Request().WithContext(ctx))

		// [Tracing log] when the request is received
		traceLogger := logger.GetTraceLogger()

		// Add traceparent if present (for OTel verification)
		recvLogEvent := traceLogger.Trace()
		if tp := c.Request().Header.Get("traceparent"); tp != "" {
			recvLogEvent.Str("traceparent", tp)
		}

		recvLogEvent.
			Str(string(logger.TraceIdKey), traceId).
			Str(string(logger.SpanIdKey), spanId).
			Str("URI", c.Request().RequestURI).
			Msg("[tracing] receive request")

		// [Tracing log] before the response is sent
		// Hooks: Before Response
		c.Response().Before(func() {
			// Log the request details
			respLogEvent := traceLogger.Trace()
			if tp := c.Request().Header.Get("traceparent"); tp != "" {
				respLogEvent.Str("traceparent", tp)
			}
			respLogEvent.
				Str(string(logger.TraceIdKey), traceId).
				Str(string(logger.SpanIdKey), spanId).
				Str("URI", c.Request().RequestURI).
				Msg("[tracing] send response")
		})

		// Call the next handler
		return next(c)
	}
}
