module github.com/middleware-labs/otel/trace

go 1.19

replace github.com/middleware-labs/otel => ../

require (
	github.com/google/go-cmp v0.5.9
	github.com/stretchr/testify v1.8.2
	github.com/middleware-labs/otel v1.15.0-rc.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/middleware-labs/otel/metric => ../metric
