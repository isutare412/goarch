package http

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

type contextKey string

const (
	contextKeyTraceID       contextKey = "ctx-key-trace-id"
	contextKeyErrorResponse contextKey = "ctx-key-error-response"
)

func injectTraceID(ctx context.Context, id trace.TraceID) context.Context {
	return context.WithValue(ctx, contextKeyTraceID, id)
}

func extractTraceID(ctx context.Context) (trace.TraceID, bool) {
	id, ok := ctx.Value(contextKeyTraceID).(trace.TraceID)
	return id, ok
}

func injectErrorResponse(ctx context.Context, err errorResponse) context.Context {
	return context.WithValue(ctx, contextKeyErrorResponse, err)
}

func extractErrorResponse(ctx context.Context) (errorResponse, bool) {
	err, ok := ctx.Value(contextKeyErrorResponse).(errorResponse)
	return err, ok
}
