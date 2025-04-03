package zlog

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/log/global"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
)

func initLogger(ctx context.Context, res *resource.Resource) (*sdklog.LoggerProvider, error) {
	exporter, err := otlploggrpc.New(ctx,
		otlploggrpc.WithEndpoint("34.122.232.161:9017"),
		otlploggrpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("creating OTLP metric exporter: %w", err)
	}
	processor := sdklog.NewBatchProcessor(exporter)

	LogProvider := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(processor),
		sdklog.WithResource(res),
	)
	global.SetLoggerProvider(LogProvider)
	return LogProvider, nil
}
