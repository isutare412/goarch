package http

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

type contextKey string

const (
	contextKeyTraceID contextKey = "ctx-key-trace-id"
)

func injectTraceID(ctx context.Context, id trace.TraceID) context.Context {
	return context.WithValue(ctx, contextKeyTraceID, id)
}

func extractTraceID(ctx context.Context) (trace.TraceID, bool) {
	id, ok := ctx.Value(contextKeyTraceID).(trace.TraceID)
	return id, ok
}
