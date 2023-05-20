module github.com/middleware-labs/otel/exporters/prometheus

go 1.19

require (
	github.com/prometheus/client_golang v1.14.0
	github.com/prometheus/client_model v0.3.0
	github.com/stretchr/testify v1.8.2
	github.com/middleware-labs/otel v1.15.0-rc.2
	github.com/middleware-labs/otel/metric v1.15.0-rc.2
	github.com/middleware-labs/otel/sdk v1.15.0-rc.2
	github.com/middleware-labs/otel/sdk/metric v0.38.0-rc.2
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/middleware-labs/otel/trace v1.15.0-rc.2 // indirect
	golang.org/x/sys v0.7.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/middleware-labs/otel => ../..

replace github.com/middleware-labs/otel/sdk => ../../sdk

replace github.com/middleware-labs/otel/sdk/metric => ../../sdk/metric

replace github.com/middleware-labs/otel/trace => ../../trace

replace github.com/middleware-labs/otel/metric => ../../metric
