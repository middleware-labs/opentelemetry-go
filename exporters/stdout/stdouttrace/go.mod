module github.com/middleware-labs/otel/exporters/stdout/stdouttrace

go 1.19

replace (
	github.com/middleware-labs/otel => ../../..
	github.com/middleware-labs/otel/sdk => ../../../sdk
)

require (
	github.com/stretchr/testify v1.8.2
	github.com/middleware-labs/otel v1.15.0-rc.2
	github.com/middleware-labs/otel/sdk v1.15.0-rc.2
	github.com/middleware-labs/otel/trace v1.15.0-rc.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/middleware-labs/otel/metric v1.15.0-rc.2 // indirect
	golang.org/x/sys v0.7.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/middleware-labs/otel/trace => ../../../trace

replace github.com/middleware-labs/otel/metric => ../../../metric
