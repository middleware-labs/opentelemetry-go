module github.com/middleware-labs/otel/example/otel-collector

go 1.19

replace (
	github.com/middleware-labs/otel => ../..
	github.com/middleware-labs/otel/sdk => ../../sdk
)

require (
	github.com/middleware-labs/otel v1.15.0-rc.2
	github.com/middleware-labs/otel/exporters/otlp/otlptrace/otlptracegrpc v1.15.0-rc.2
	github.com/middleware-labs/otel/sdk v1.15.0-rc.2
	github.com/middleware-labs/otel/trace v1.15.0-rc.2
	google.golang.org/grpc v1.54.0
)

require (
	github.com/cenkalti/backoff/v4 v4.2.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.0 // indirect
	github.com/middleware-labs/otel/exporters/otlp/internal/retry v1.15.0-rc.2 // indirect
	github.com/middleware-labs/otel/exporters/otlp/otlptrace v1.15.0-rc.2 // indirect
	github.com/middleware-labs/otel/metric v1.15.0-rc.2 // indirect
	go.opentelemetry.io/proto/otlp v0.19.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

replace github.com/middleware-labs/otel/trace => ../../trace

replace github.com/middleware-labs/otel/exporters/otlp/otlptrace => ../../exporters/otlp/otlptrace

replace github.com/middleware-labs/otel/exporters/otlp/otlptrace/otlptracegrpc => ../../exporters/otlp/otlptrace/otlptracegrpc

replace github.com/middleware-labs/otel/exporters/otlp/internal/retry => ../../exporters/otlp/internal/retry

replace github.com/middleware-labs/otel/metric => ../../metric
