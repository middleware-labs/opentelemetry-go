module github.com/middleware-labs/otel/example/opencensus

go 1.19

replace (
	github.com/middleware-labs/otel => ../..
	github.com/middleware-labs/otel/bridge/opencensus => ../../bridge/opencensus
	github.com/middleware-labs/otel/sdk => ../../sdk
)

require (
	go.opencensus.io v0.24.0
	github.com/middleware-labs/otel v1.15.0-rc.2
	github.com/middleware-labs/otel/bridge/opencensus v0.38.0-rc.2
	github.com/middleware-labs/otel/exporters/stdout/stdoutmetric v0.38.0-rc.2
	github.com/middleware-labs/otel/exporters/stdout/stdouttrace v1.15.0-rc.2
	github.com/middleware-labs/otel/sdk v1.15.0-rc.2
	github.com/middleware-labs/otel/sdk/metric v0.38.0-rc.2
)

require (
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/middleware-labs/otel/metric v1.15.0-rc.2 // indirect
	github.com/middleware-labs/otel/trace v1.15.0-rc.2 // indirect
	golang.org/x/sys v0.7.0 // indirect
)

replace github.com/middleware-labs/otel/metric => ../../metric

replace github.com/middleware-labs/otel/sdk/metric => ../../sdk/metric

replace github.com/middleware-labs/otel/trace => ../../trace

replace github.com/middleware-labs/otel/exporters/stdout/stdoutmetric => ../../exporters/stdout/stdoutmetric

replace github.com/middleware-labs/otel/exporters/stdout/stdouttrace => ../../exporters/stdout/stdouttrace
