package observability

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
	"os"
	"sync"
)

var (
	onceSlog   sync.Once
	slogLogger *slog.Logger
)

func initSlog(serviceName, version string) *slog.Logger {
	// Log initialization
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			return a
		},
	}).WithAttrs([]slog.Attr{
		slog.String("service", serviceName),
		slog.String("with-release", version),
	})
	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	return logger
}

func NewSlog(serviceName, version string) {
	onceSlog.Do(func() {
		log := initSlog(serviceName, version)
		h := otelHandler{h: log.Handler()}
		l := slog.New(h)
		slogLogger = l
	})

	slog.SetDefault(slogLogger)
}

type otelHandler struct{ h slog.Handler }

func (s otelHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return true /* support all logging levels*/
}

func (s otelHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return otelHandler{h: s.h.WithAttrs(attrs)}
}

func (s otelHandler) WithGroup(name string) slog.Handler {
	return otelHandler{h: s.h.WithGroup(name)}
}

func (s otelHandler) Handle(ctx context.Context, r slog.Record) (err error) {

	if ctx == nil {
		return s.h.Handle(ctx, r)
	}

	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return s.h.Handle(ctx, r)
	}

	{
		sCtx := span.SpanContext()
		attrs := make([]slog.Attr, 0)
		if sCtx.HasTraceID() {
			attrs = append(attrs,
				slog.Attr{Key: "traceId", Value: slog.StringValue(sCtx.TraceID().String())},
			)
		}
		if sCtx.HasSpanID() {
			attrs = append(attrs,
				slog.Attr{Key: "spanId", Value: slog.StringValue(sCtx.SpanID().String())},
			)
		}
		if len(attrs) > 0 {
			r.AddAttrs(attrs...)
		}
	}

	{
		attrs := make([]attribute.KeyValue, 0)

		logSeverityKey := attribute.Key("log.severity")
		logMessageKey := attribute.Key("log.message")
		attrs = append(attrs, logSeverityKey.String(r.Level.String()))
		attrs = append(attrs, logMessageKey.String(r.Message))

		r.Attrs(func(a slog.Attr) bool {
			attrs = append(attrs,
				attribute.KeyValue{
					Key:   attribute.Key(a.Key),
					Value: attribute.StringValue(a.Value.String()),
				},
			)
			return true
		})

		span.AddEvent(fmt.Sprintf("%s: %s", r.Level, r.Message), trace.WithAttributes(attrs...))
		if r.Level >= slog.LevelError {
			span.SetStatus(codes.Error, r.Message)
		}
	}

	return s.h.Handle(ctx, r)
}
