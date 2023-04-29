package tracing

import (
	"context"
	"fmt"
	"sync/atomic"

	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

var (
	globalTracerProvider atomic.Pointer[sdktrace.TracerProvider]
)

func Init(cfg Config) {
	if !cfg.Enabled {
		return
	}

	exp, err := newJaegerExporter(cfg)
	if err != nil {
		panic(fmt.Errorf("creating Jaeger exporter: %w", err))
	}

	prv := newTracerProvider(cfg, exp)
	globalTracerProvider.Store(prv)

	otel.SetTracerProvider(prv)
	otel.SetTextMapPropagator(b3.New(b3.WithInjectEncoding(b3.B3MultipleHeader)))
}

func Shutdown() error {
	prv := globalTracerProvider.Load()
	if prv == nil {
		return nil
	}
	return prv.Shutdown(context.Background())
}

func newJaegerExporter(cfg Config) (*jaeger.Exporter, error) {
	return jaeger.New(jaeger.WithCollectorEndpoint(
		jaeger.WithEndpoint(cfg.JaegerCollectorEndpoint),
	))
}

func newTracerProvider(cfg Config, exp sdktrace.SpanExporter) *sdktrace.TracerProvider {
	rsc := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(cfg.ServiceName),
		semconv.DeploymentEnvironment(cfg.Environment),
	)

	prv := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.AlwaysSample())),
		sdktrace.WithResource(rsc),
	)
	return prv
}
