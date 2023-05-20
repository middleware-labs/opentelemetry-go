# OpenTelemetry Exporters

Once the OpenTelemetry SDK has created and processed telemetry, it needs to be exported.
This package contains exporters for this purpose.

## Exporter Packages

The following exporter packages are provided with the following OpenTelemetry signal support.

| Exporter Package                                                                | Metrics | Traces |
| :-----------------------------------------------------------------------------: | :-----: | :----: |
| [github.com/middleware-labs/otel/exporters/jaeger](./jaeger)                           |         | ✓      |
| [github.com/middleware-labs/otel/exporters/otlp/otlpmetric](./otlp/otlpmetric)         | ✓       |        |
| [github.com/middleware-labs/otel/exporters/otlp/otlptrace](./otlp/otlptrace)           |         | ✓      |
| [github.com/middleware-labs/otel/exporters/prometheus](./prometheus)                   | ✓       |        |
| [github.com/middleware-labs/otel/exporters/stdout/stdoutmetric](./stdout/stdoutmetric) | ✓       |        |
| [github.com/middleware-labs/otel/exporters/stdout/stdouttrace](./stdout/stdouttrace)   |         | ✓      |
| [github.com/middleware-labs/otel/exporters/zipkin](./zipkin)                           |         | ✓      |

See the [OpenTelemetry registry] for 3rd-part exporters compatible with this project.

[OpenTelemetry registry]: https://opentelemetry.io/registry/?language=go&component=exporter
