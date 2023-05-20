module github.com/middleware-labs/otel/bridge/opentracing/test

go 1.19

replace github.com/middleware-labs/otel => ../../..

replace github.com/middleware-labs/otel/bridge/opentracing => ../

replace github.com/middleware-labs/otel/trace => ../../../trace

require (
	github.com/opentracing-contrib/go-grpc v0.0.0-20210225150812-73cb765af46e
	github.com/opentracing/opentracing-go v1.2.0
	github.com/stretchr/testify v1.8.2
	github.com/middleware-labs/otel v1.15.0-rc.2
	github.com/middleware-labs/otel/bridge/opentracing v1.15.0-rc.2
	google.golang.org/grpc v1.54.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/middleware-labs/otel/metric v1.15.0-rc.2 // indirect
	github.com/middleware-labs/otel/trace v1.15.0-rc.2 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/middleware-labs/otel/metric => ../../../metric
