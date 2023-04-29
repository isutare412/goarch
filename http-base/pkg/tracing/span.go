package tracing

import (
	"context"
	"net/http"
	"path"
	"runtime"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

const (
	defaultTracer = "default"
)

// AutoSpan generates a span with caller function name. If you need to use
// custom function name, use ManualSpan instead.
func AutoSpan(ctx context.Context) (context.Context, trace.Span) {
	return otel.Tracer(defaultTracer).Start(ctx, funcName(2))
}

// ManualSpan generates a span with given name.
func ManualSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	return otel.Tracer(defaultTracer).Start(ctx, name)
}

// AutoError record error to span with caller function name. If you need to use
// custom function name, use ManualError instead.
func AutoError(span trace.Span, err error) {
	recordError(span, err, funcName(2))
}

// ManualError record error to span with given name.
func ManualError(span trace.Span, err error, name string) {
	recordError(span, err, name)
}

func recordError(span trace.Span, err error, name string) {
	span.SetStatus(codes.Error, name)
	span.RecordError(err)
}

func ExtractIDs(span trace.Span) (trace.TraceID, trace.SpanID) {
	sc := span.SpanContext()
	return sc.TraceID(), sc.SpanID()
}

func CtxFromHTTPHeader(ctx context.Context, h http.Header) context.Context {
	return otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(h))
}

func CtxFromMap(ctx context.Context, m map[string]string) context.Context {
	return otel.GetTextMapPropagator().Extract(ctx, propagation.MapCarrier(m))
}

func SerializeIntoMap(ctx context.Context) map[string]string {
	// Capacity of map is 5 because we are using B3MultipleHeader for
	// propagation.
	m := make(map[string]string, 5)
	otel.GetTextMapPropagator().Inject(ctx, propagation.MapCarrier(m))
	return m
}

func funcName(skip int) string {
	name := "unknownFunc"

	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		return name
	}

	caller := runtime.FuncForPC(pc)
	if caller == nil {
		return name
	}
	name = path.Base(caller.Name())

	return name
}
