package http

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	"github.com/isutare412/goarch/http-base/pkg/log"
	"github.com/isutare412/goarch/http-base/pkg/tracing"
)

func startTrace(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := tracing.CtxFromHTTPHeader(r.Context(), r.Header)
		ctx, span := tracing.AutoSpan(ctx)
		defer span.End()

		traceID, _ := tracing.ExtractIDs(span)
		ctx = injectTraceID(ctx, traceID)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func requestLogger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		beforeServing := time.Now()
		next.ServeHTTP(ww, r)
		elapsedTime := time.Since(beforeServing)

		var contentType string
		if ct := r.Header.Get("Content-Type"); ct != "" {
			contentType = ct
		}

		var statusCode = http.StatusOK
		if sc := ww.Status(); sc != 0 {
			statusCode = sc
		}

		var traceID string
		if id, ok := extractTraceID(r.Context()); ok {
			traceID = id.String()
		}

		log.A().
			With(
				zap.String("method", r.Method),
				zap.String("url", r.URL.String()),
				zap.String("addr", r.RemoteAddr),
				zap.String("proto", r.Proto),
				zap.String("traceID", traceID),
				zap.String("contentType", contentType),
				zap.Int64("contentLength", r.ContentLength),
				zap.String("userAgent", r.UserAgent()),
				zap.Int("status", statusCode),
				zap.Int("bodyBytes", ww.BytesWritten()),
				zap.Duration("elapsed", elapsedTime),
			).Info()
	}

	return http.HandlerFunc(fn)
}

func recoverPanic(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.WithOperation("recoverHTTPPanic").With(zap.Stack("stackTrace")).Errorf("panicked: %v", r)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
