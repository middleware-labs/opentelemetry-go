module github.com/middleware-labs/otel/example/jaeger

go 1.19

replace (
	github.com/middleware-labs/otel => ../..
	github.com/middleware-labs/otel/exporters/jaeger => ../../exporters/jaeger
	github.com/middleware-labs/otel/sdk => ../../sdk
)

require (
	github.com/middleware-labs/otel v1.15.0-rc.2
	github.com/middleware-labs/otel/exporters/jaeger v1.15.0-rc.2
	github.com/middleware-labs/otel/sdk v1.15.0-rc.2
)

require (
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/middleware-labs/otel/metric v1.15.0-rc.2 // indirect
	github.com/middleware-labs/otel/trace v1.15.0-rc.2 // indirect
	golang.org/x/sys v0.7.0 // indirect
)

replace github.com/middleware-labs/otel/trace => ../../trace

replace github.com/middleware-labs/otel/metric => ../../metric
